package consumer

import "github.com/segmentio/kafka-go"

func New(server, topic string) kafka.Reader {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{server},
		Topic:   topic,
	})
	return *reader
}
