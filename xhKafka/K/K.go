package K

import (
	"github.com/segmentio/kafka-go"
	"github.com/xh-dev-go/xhUtils/xhKafka/KProducer"
	"github.com/xh-dev-go/xhUtils/xhKafka/KWildcard"
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

func (k *K) Reader(wildcard KWildcard.KafkaWildcard) (chan kafka.Message, chan error, kafka.Reader){
	reader := NewReader(k.Server, k.Topic)
	matchChan, errChan := wildcard.BindWildcard(reader, GetMessageChan(), GetErrChan())
	return matchChan, errChan, reader
}

func (k *K) Producer(key string) KProducer.XhKafkaProducer{
	return KProducer.New(k.Server, k.Topic, key)
}

func NewReader(server, topic string) kafka.Reader {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{server},
		Topic:   topic,
	})
	return *reader
}

func GetMessageChan() chan kafka.Message{
	return make(chan kafka.Message)
}
func GetErrChan() chan error{
	return make(chan error)
}

