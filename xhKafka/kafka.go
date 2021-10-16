package xhKafka

import (
	"github.com/segmentio/kafka-go"
	"strings"
)

type KafkaHeaders []kafka.Header

func (i *KafkaHeaders) ToKafkaHeaders() []kafka.Header {
	return *i
}

func (i *KafkaHeaders) String() string {
	var str = ""

	for _, item := range *i {
		str += item.Key
		str += "="
		str += string(item.Value)
		str += ", "
	}
	if len(str) > 0 {
		return str[:len(str)-2]
	} else {
		return str
	}
}

func (i *KafkaHeaders) Set(value string) error {
	vs := strings.Split(value, "=")
	if len(vs) != 2 {
		panic("Header not with correct format: " + value)
	} else {
		*i = append(*i, kafka.Header{
			Key:   vs[0],
			Value: []byte(vs[1]),
		})
	}
	return nil
}
