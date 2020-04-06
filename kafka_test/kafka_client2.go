package main

import (
	"context"
	"github.com/segmentio/kafka-go"
	"time"
)

func main() {
	topic := "test-topic"
	partition := 0

	conn, _ := kafka.DialLeader(context.Background(), "tcp", "kafka.dev.9rum.cc:30100", topic, partition)

	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	conn.WriteMessages(
		kafka.Message{Value: []byte("kafka-go one!")},
		kafka.Message{Value: []byte("kafka-go two!")},
		kafka.Message{Value: []byte("kafka-go three!")},
	)

	conn.Close()
}
