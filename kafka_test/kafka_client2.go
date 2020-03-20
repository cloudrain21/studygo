package main

import (
	"context"
	"github.com/segmentio/kafka-go"
	"time"
)

func main() {
	topic := "rain-topic"
	partition := 0

	conn, _ := kafka.DialLeader(context.Background(), "tcp", "kk001:9092", topic, partition)

	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	conn.WriteMessages(
		kafka.Message{Value: []byte("kafka-go one!")},
		kafka.Message{Value: []byte("kafka-go two!")},
		kafka.Message{Value: []byte("kafka-go three!")},
	)

	conn.Close()
}
