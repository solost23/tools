package kafka

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"

	"github.com/Shopify/sarama"
)

type ConsumeService interface {
	Consume([]byte) error
}

// Consumer represents a Sarama consumer group consumer
type Consumer struct {
	ready         chan bool
	client        sarama.Client
	consumeClient ConsumeService
	qid           string
	group         string
	topics        string
}

func NewConsumerClient(brokers []string) (sarama.Client, error) {
	assignor := "range"
	config := sarama.NewConfig()
	config.Version = sarama.V0_10_2_1
	switch assignor {
	case "sticky":
		config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategySticky
	case "roundrobin":
		config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRoundRobin
	case "range":
		config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRange
	default:
		log.Panicf("Unrecognized consumer group partition assignor: %s", assignor)
	}
	client, err := sarama.NewClient(brokers, config)
	return client, err
}

func NewConsumer(client sarama.Client, topics string, group string, consume ConsumeService) *Consumer {
	return &Consumer{
		ready:         make(chan bool),
		client:        client,
		consumeClient: consume,
		group:         group,
		topics:        topics,
	}
}

func (consumer *Consumer) Consume() {
	ctx, cancel := context.WithCancel(context.Background())
	consumerClient, err := sarama.NewConsumerGroupFromClient(consumer.group, consumer.client)
	if err != nil {
		log.Panicf("Error creating consumer group client: %v", err)
	}
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			// `Consume` should be called inside an infinite loop, when a
			// server-side rebalance happens, the consumer session will need to be
			// recreated to get the new claims
			if err := consumerClient.Consume(ctx, strings.Split(consumer.topics, ","), consumer); err != nil {
				log.Panicf("Error from consumer: %v", err)
			}
			// check if context was cancelled, signaling that the consumer should stop
			if ctx.Err() != nil {
				return
			}
			consumer.ready = make(chan bool)
		}
	}()
	<-consumer.ready
	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-ctx.Done():
		log.Println(ctx, "terminating: context cancelled")
	case <-sigterm:
		log.Println(ctx, "terminating: via signal")
	}
	cancel()
	wg.Wait()
	if err = consumerClient.Close(); err != nil {
		log.Println("consumer", fmt.Sprintf("Error closing client: %v", err))
	}
	log.Println("consumer", "exit")
}

// Setup is run at the beginning of a new session, before ConsumeClaim
func (consumer *Consumer) Setup(sarama.ConsumerGroupSession) error {
	// Mark the consumer as ready
	close(consumer.ready)
	return nil
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
func (consumer *Consumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
func (consumer *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		log.Println(consumer.qid, fmt.Sprintf("Message claimed: value = %s, timestamp = %v, topic = %s", string(message.Value), message.Timestamp, message.Topic))
		err := consumer.consumeClient.Consume(message.Value)
		if err != nil {
			// 发生错误打印log
			log.Println(consumer.qid, err.Error())
		}
		session.MarkMessage(message, "")
	}
	return nil
}
