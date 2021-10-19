package kafkaWrapper

import (
	"github.com/segmentio/kafka-go"
	"github.com/xh-dev-go/xhUtils/xhKafka/channel"
	"github.com/xh-dev-go/xhUtils/xhKafka/consumer"
	"github.com/xh-dev-go/xhUtils/xhKafka/producer"
	"github.com/xh-dev-go/xhUtils/xhKafka/wildcard"
)

type K struct {
	Server string
	Topic string
}

func New(server string, topic string) K {
	return K{
		Server: server,
		Topic: topic,
	}
}

func (k *K) read(wildcard wildcard.KafkaWildcard) (chan kafka.Message, chan error){
	reader := consumer.New(k.Server, k.Topic)
	return wildcard.BindWildcard(reader, channel.GetMessageChan(), channel.GetErrChan())
}

func (k *K) producer(key string) producer.XhKafkaProducer{
	return producer.New(k.Server, k.Topic, key)
}

