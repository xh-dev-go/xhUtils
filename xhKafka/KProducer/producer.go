package KProducer

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"github.com/xh-dev-go/xhUtils/xhKafka/KHeader"
)

type XhKafkaProducer struct {
	Topic string
	Key string
	writer *kafka.Writer
}

func New(server, topic, key string) XhKafkaProducer {
	w := &kafka.Writer{
		Addr: kafka.TCP(server),
		// NOTE: When Topic is not defined here, each Message must define it instead.
		Balancer: &kafka.LeastBytes{},
	}
	return XhKafkaProducer{
		Topic: topic,
		Key: key,
		writer: w,
	}
}
func (xhKafkaProducer *XhKafkaProducer) ToWriter() *kafka.Writer  {
	return xhKafkaProducer.writer
}
func (xhKafkaProducer *XhKafkaProducer) SimpleSend(msg kafka.Message) error {
	println(fmt.Sprintf("%s: %s", string(msg.Key),msg.Value))
	return xhKafkaProducer.ToWriter().WriteMessages(
		context.Background(),
		msg,
	)
}
func (xhKafkaProducer *XhKafkaProducer) Message(msg string, headers KHeader.KafkaHeaders) kafka.Message {
	return kafka.Message{
		Key: []byte(xhKafkaProducer.Key),
		Topic: xhKafkaProducer.Topic,
		Value: []byte(msg),
		Headers: headers,
	}
}
