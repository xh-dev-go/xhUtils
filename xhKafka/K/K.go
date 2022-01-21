package K

import (
	"context"
	"errors"
	"github.com/google/uuid"
	kafka "github.com/segmentio/kafka-go"
	"github.com/xh-dev-go/xhUtils/xhKafka/KHeader"
	"github.com/xh-dev-go/xhUtils/xhKafka/KKey"
	"github.com/xh-dev-go/xhUtils/xhKafka/KProducer"
	"github.com/xh-dev-go/xhUtils/xhKafka/KWildcard"
	"time"
)

type K struct {
	Server string
	Topic  string
}

func New(server string, topic string) *K {
	return &K{
		Server: server,
		Topic:  topic,
	}
}

func (k *K) ReadFromNow(wildcard KWildcard.KafkaWildcard) (chan kafka.Message, chan error, *kafka.Reader, error) {
	return k.ReadSinceTime(wildcard, time.Now())
}

func (k *K) ReadSinceTime(wildcard KWildcard.KafkaWildcard, time time.Time) (chan kafka.Message, chan error, *kafka.Reader, error) {
	reader := NewReader(k.Server, k.Topic)
	fun := func(reader2 *kafka.Reader) error {
		err := reader2.SetOffsetAt(context.Background(), time)
		if err != nil {
			return err
		}
		return nil
	}
	matchChan, errChan, err := wildcard.BindWildcard(&reader, GetMessageChan(), GetErrChan(), fun)
	return matchChan, errChan, &reader, err
}

func (k *K) ReadAtStart(wildcard KWildcard.KafkaWildcard) (chan kafka.Message, chan error, *kafka.Reader, error) {
	return k.ReaderAtOffset(wildcard, 0)
}

func (k *K) ReaderAtOffset(wildcard KWildcard.KafkaWildcard, offset int64) (chan kafka.Message, chan error, *kafka.Reader, error) {
	reader := NewReader(k.Server, k.Topic)
	fun := func(reader2 *kafka.Reader) error {
		err := reader2.SetOffset(offset)
		if err != nil {
			return err
		}
		return nil
	}
	matchChan, errChan, err := wildcard.BindWildcard(&reader, GetMessageChan(), GetErrChan(), fun)
	return matchChan, errChan, &reader, err
}

func (k *K) Producer(key string) KProducer.XhKafkaProducer {
	return KProducer.New(k.Server, k.Topic, key)
}

func NewReader(server, topic string) kafka.Reader {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{server},
		Topic:   topic,
	})
	return *reader
}

func GetMessageChan() chan kafka.Message {
	return make(chan kafka.Message)
}
func GetErrChan() chan error {
	return make(chan error)
}


func (k *K) EmptyMessage(key *KKey.KKey) kafka.Message {
	return kafka.Message{
		Key:     []byte(key.ToKey()),
		Topic:   k.Topic,
		Value:   []byte(""),
		Headers: KHeader.BasicHeader(),
	}
}
func (k *K) SimpleMessage(key *KKey.KKey, msg string) kafka.Message {
	return kafka.Message{
		Key:     []byte(key.ToKey()),
		Topic:   k.Topic,
		Value:   []byte(msg),
		Headers: KHeader.BasicHeader(),
	}
}
func (k *K) SimpleSend(message kafka.Message) error {
	producer := k.Producer(message.Topic)
	return producer.SimpleSend(message)
}

type Response struct {
	Msg kafka.Message
	Err error
}

func (resp Response) HasError() bool {
	return resp.Err != nil
}

func (resp Response) NoError() bool {
	return !resp.HasError()
}
func (resp Response) IsTimeout() bool {
	if resp.NoError() {
		return false
	} else {
		return resp.Err == RequestTimeError
	}
}

var RequestTimeError = errors.New("timeout")
var FailToCreateUuid = errors.New("fail to create uuid")

func (k K) Request(key *KKey.KKey, msg kafka.Message, timeout time.Duration) Response {
	if id, err := uuid.NewUUID(); err != nil {
		request := key.Request().Push(id.String())
		msg.Key = []byte(request.ToKey())

		response := key.Response().Push(id.String())
		msgChan, errChan, _, err := k.ReadFromNow(response.ToWildcard())
		if err != nil {
			return Response{
				Err: err,
			}
		}
		timeoutChan := make(chan bool)
		go func() {
			time.Sleep(timeout)
			timeoutChan <- true
		}()

		err = k.SimpleSend(msg)
		if err != nil {
			return Response{Err: err}
		}

		select {
		case <-timeoutChan:
			return Response{Err: RequestTimeError}
		case msg := <-msgChan:
			return Response{Msg: msg}
		case err := <-errChan:
			return Response{Err: err}
		}
	} else {
		return Response{Err: FailToCreateUuid}
	}
}
