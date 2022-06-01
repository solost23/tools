package kafka

import (
	"testing"
	"time"
)

func TestNewConsumer(t *testing.T) {
	a := make(chan bool)
	go func() {
		time.Sleep(1 * time.Second)
		a = make(chan bool)
	}()
	<-a
	t.Log("success")
}
