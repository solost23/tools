package kafka

import (
	"fmt"
	"testing"
)

func TestProducer_Publish(t *testing.T) {
	topic := "chinese"
	client, err := NewProducerClient([]string{"192.168.137.1:9092"})
	if err != nil {
		panic(err.Error())
	}
	producer := NewProducer(client, topic)
	for i := 0; i < 1000; i++ {
		p, offset, err := producer.Publish([]byte(fmt.Sprintf("message:%d", i)))
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(p, offset)
	}
}
