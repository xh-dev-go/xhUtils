package producer

import (
	"context"
	"github.com/segmentio/kafka-go"
)

type XhKafkaProducer kafka.Writer

func New(server string) XhKafkaProducer {
	w := &kafka.Writer{
		Addr: kafka.TCP(server),
		// NOTE: When Topic is not defined here, each Message must define it instead.
		Balancer: &kafka.LeastBytes{},
	}
	return *(*XhKafkaProducer)(w)
}
func (w *XhKafkaProducer) ToWriter() *kafka.Writer  {
	return (*kafka.Writer)(w)
}
func (w *XhKafkaProducer) SimpleSend(msg kafka.Message) error {
	return w.ToWriter().WriteMessages(
		context.Background(),
		msg,
	)
}
