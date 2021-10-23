package K

import (
	"context"
	kafka "github.com/segmentio/kafka-go"
	"github.com/xh-dev-go/xhUtils/xhKafka/KProducer"
	"github.com/xh-dev-go/xhUtils/xhKafka/KWildcard"
	"time"
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

func (k *K) ReadFromNow(wildcard KWildcard.KafkaWildcard) (chan kafka.Message, chan error, *kafka.Reader, error){
	return k.ReadSinceTime(wildcard, time.Now())
}

func (k *K) ReadSinceTime(wildcard KWildcard.KafkaWildcard, time time.Time) (chan kafka.Message, chan error, *kafka.Reader, error){
	reader := NewReader(k.Server, k.Topic)
	fun := func (reader2 *kafka.Reader) error {
		err := reader2.SetOffsetAt(context.Background(), time)
		if err != nil {
			return err
		}
		return nil
	}
	matchChan, errChan, err := wildcard.BindWildcard(&reader, GetMessageChan(), GetErrChan(), fun)
	return matchChan, errChan, &reader, err
}

func (k *K) ReadAtStart(wildcard KWildcard.KafkaWildcard) (chan kafka.Message, chan error, *kafka.Reader, error){
	return k.ReaderAtOffset(wildcard, 0)
}

func (k *K) ReaderAtOffset(wildcard KWildcard.KafkaWildcard, offset int64) (chan kafka.Message, chan error, *kafka.Reader, error){
	reader := NewReader(k.Server, k.Topic)
	fun := func (reader2 *kafka.Reader) error {
		err := reader2.SetOffset(offset)
		if err != nil {
			return err
		}
		return nil
	}
	matchChan, errChan, err := wildcard.BindWildcard(&reader, GetMessageChan(), GetErrChan(), fun)
	return matchChan, errChan, &reader, err
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

