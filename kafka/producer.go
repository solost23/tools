package kafka

import "github.com/Shopify/sarama"

type Producer struct {
	client sarama.SyncProducer
	topic  string
}

func NewProducerClient(addrs []string) (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll // 赋值为 -1 意味着producer在follower副本确认接收到数据后才算一次发送完成
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true
	config.Version = sarama.V0_10_2_1
	client, err := sarama.NewSyncProducer(addrs, config)
	return client, err
}

func NewProducer(client sarama.SyncProducer, topic string) *Producer {
	return &Producer{
		client: client,
		topic:  topic,
	}
}

func (p *Producer) Publish(msg []byte) (pid int32, offset int64, err error) {
	producerMessage := &sarama.ProducerMessage{}
	producerMessage.Topic = p.topic
	producerMessage.Value = sarama.ByteEncoder(msg)
	pid, offset, err = p.client.SendMessage(producerMessage)
	if err != nil {
		return
	}
	return
}

func (p *Producer) Close() {
	p.client.Close()
}
