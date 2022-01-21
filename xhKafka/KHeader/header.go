package KHeader

import (
	"github.com/segmentio/kafka-go"
	"strings"
	"time"
)

type KafkaHeaders []kafka.Header

func BasicHeader() KafkaHeaders{
	var header = FromKafkaHeader([]kafka.Header{})
	header = header.Add("Date", time.Now().Format(time.RFC3339))
	return header
}
func FromKafkaHeader(headers []kafka.Header) KafkaHeaders {
	return headers
}

func (i *KafkaHeaders) Find1(name string) string {
	for _, item := range i.ToKafkaHeaders() {
		if item.Key == name {
			return string(item.Value)
		}
	}
	return ""
}

func (i *KafkaHeaders) Find2(name1, name2 string) (string, string) {
	var v1, v2 string
	for _, item := range i.ToKafkaHeaders() {
		if item.Key == name1 {
			v1 = string(item.Value)
		}
		if item.Key == name2 {
			v2 = string(item.Value)
		}
		if v1 != "" && v2 != "" {
			return v1, v2
		}
	}
	return v1, v2
}

func (i *KafkaHeaders) Find3(name1, name2, name3 string) (string, string, string) {
	var v1, v2, v3 string
	for _, item := range i.ToKafkaHeaders() {
		if item.Key == name1 {
			v1 = string(item.Value)
		}
		if item.Key == name2 {
			v2 = string(item.Value)
		}
		if item.Key == name3 {
			v3 = string(item.Value)
		}
		if v1 != "" && v2 != "" && v3 != "" {
			return v1, v2, v3
		}
	}
	return v1, v2, v3
}

func (i *KafkaHeaders) Find4(name1, name2, name3, name4 string) (string, string, string, string) {
	var v1, v2, v3, v4 string
	for _, item := range i.ToKafkaHeaders() {
		if item.Key == name1 {
			v1 = string(item.Value)
		}
		if item.Key == name2 {
			v2 = string(item.Value)
		}
		if item.Key == name3 {
			v3 = string(item.Value)
		}
		if item.Key == name4 {
			v4 = string(item.Value)
		}
		if v1 != "" && v2 != "" && v3 != "" && v4 != "" {
			return v1, v2, v3, v4
		}
	}
	return v1, v2, v3, v4
}

func (i *KafkaHeaders) Find5(name1, name2, name3, name4, name5 string) (string, string, string, string, string) {
	var v1, v2, v3, v4, v5 string
	for _, item := range i.ToKafkaHeaders() {
		if item.Key == name1 {
			v1 = string(item.Value)
		}
		if item.Key == name2 {
			v2 = string(item.Value)
		}
		if item.Key == name3 {
			v3 = string(item.Value)
		}
		if item.Key == name4 {
			v4 = string(item.Value)
		}
		if item.Key == name5 {
			v5 = string(item.Value)
		}
		if v1 != "" && v2 != "" && v3 != "" && v4 != "" && v5 != "" {
			return v1, v2, v3, v4, v5
		}
	}
	return v1, v2, v3, v4, v5
}

func (i *KafkaHeaders) Find6(name1, name2, name3, name4, name5, name6 string) (string, string, string, string, string, string) {
	var v1, v2, v3, v4, v5, v6 string
	for _, item := range i.ToKafkaHeaders() {
		if item.Key == name1 {
			v1 = string(item.Value)
		}
		if item.Key == name2 {
			v2 = string(item.Value)
		}
		if item.Key == name3 {
			v3 = string(item.Value)
		}
		if item.Key == name4 {
			v4 = string(item.Value)
		}
		if item.Key == name5 {
			v5 = string(item.Value)
		}
		if item.Key == name6 {
			v6 = string(item.Value)
		}
		if v1 != "" && v2 != "" && v3 != "" && v4 != "" && v5 != "" && v6 != "" {
			return v1, v2, v3, v4, v5, v6
		}
	}
	return v1, v2, v3, v4, v5, v6
}

func (i *KafkaHeaders) Find7(name1, name2, name3, name4, name5, name6, name7 string) (string, string, string, string, string, string, string) {
	var v1, v2, v3, v4, v5, v6, v7 string
	for _, item := range i.ToKafkaHeaders() {
		if item.Key == name1 {
			v1 = string(item.Value)
		}
		if item.Key == name2 {
			v2 = string(item.Value)
		}
		if item.Key == name3 {
			v3 = string(item.Value)
		}
		if item.Key == name4 {
			v4 = string(item.Value)
		}
		if item.Key == name5 {
			v5 = string(item.Value)
		}
		if item.Key == name6 {
			v6 = string(item.Value)
		}
		if item.Key == name7 {
			v7 = string(item.Value)
		}
		if v1 != "" && v2 != "" && v3 != "" && v4 != "" && v5 != "" && v6 != "" && v7 != "" {
			return v1, v2, v3, v4, v5, v6, v7
		}
	}
	return v1, v2, v3, v4, v5, v6, v7
}

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

func (i *KafkaHeaders) Add(key, value string) KafkaHeaders {
	t := FromKafkaHeader(append(*i, kafka.Header{
		Key:   key,
		Value: []byte(value),
	}))
	return t
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
