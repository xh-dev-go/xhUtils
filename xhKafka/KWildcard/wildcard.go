package KWildcard

import (
	"context"
	"github.com/segmentio/kafka-go"
	"github.com/xh-dev-go/xhUtils/flagUtils"
	"io"
	"strings"
)

type KafkaWildcard struct {
	levels []string
}

func New(path string) KafkaWildcard {
	return KafkaWildcard{levels: strings.Split(strings.Trim(path, "/"), "/")}
}

func (wildcard KafkaWildcard) Len() int {
	return len(wildcard.levels)
}

func (wildcard KafkaWildcard) Last() string {
	return wildcard.levels[wildcard.Len()-1]
}

func getMessageChan() chan kafka.Message{
	return make(chan kafka.Message)
}
func getErrChan() chan error{
	return make(chan error)
}
func (wildcard KafkaWildcard) Match(key string) bool {
	splintedKey := strings.Split(strings.Trim(key, "/"), "/")
	keyLen := len(splintedKey)

	if keyLen < wildcard.Len() {
		return false
	}

	for index, level := range wildcard.levels {
		if index >= keyLen {
			break
		}

		switch level {
		case "*":
			continue
		case ">":
			return true
		default:
			if level == splintedKey[index] {
				continue
			} else {
				return false
			}
		}

	}

	if wildcard.Last() == "*" && wildcard.Len() != keyLen {
		return false
	} else {
		return true
	}

}
func (wildcard KafkaWildcard) BindWildcard(reader *kafka.Reader, matchChan chan kafka.Message, errChan chan error, extraSettings func (reader2 *kafka.Reader) error )  (chan kafka.Message, chan error, error){
	err := extraSettings(reader)
	if err != nil {
		return nil, nil, err
	}
	go func() {
		defer func(reader *kafka.Reader) {
			err := reader.Close()
			if err != nil {

			}
		}(reader)
		for {
			if m, err := reader.ReadMessage(context.Background()); err != nil {
				if err == io.EOF {
					println("Closed")
					break
				} else {
					errChan <- err
					continue
				}
			} else {
				if wildcard.Match(string(m.Key)) {
					matchChan <- m
				}
			}
		}
	}()
	return matchChan, errChan, nil
}

func (wildcard KafkaWildcard) ReverseKey1(key string, index int) (bool, string){
	if wildcard.Match(key) {
		key := strings.Split(strings.Trim(key, "/"), "/")
		if len(key) <= index {
			return false, flagUtils.StrEmpty
		}
		return true, key[index]
	} else {
		return false, flagUtils.StrEmpty
	}
}

func (wildcard KafkaWildcard) ReverseKey2(key string, index1, index2 int) (bool, string, string){
	if wildcard.Match(key) {
		key := strings.Split(strings.Trim(key, "/"), "/")
		if l:=len(key); l <= index1 || l <= index2 {
			return false, flagUtils.StrEmpty, flagUtils.StrEmpty
		}
		return true, key[index1], key[index2]
	} else {
		return false, flagUtils.StrEmpty, flagUtils.StrEmpty
	}
}

func (wildcard KafkaWildcard) ReverseKey3(key string, index1, index2, index3 int) (bool, string, string, string){
	if wildcard.Match(key) {
		key := strings.Split(strings.Trim(key, "/"), "/")
		if l:=len(key); l <= index1 || l <= index2 || l <= index3{
			return false, flagUtils.StrEmpty, flagUtils.StrEmpty, flagUtils.StrEmpty
		}
		return true, key[index1], key[index2], key[index3]
	} else {
		return false, flagUtils.StrEmpty, flagUtils.StrEmpty, flagUtils.StrEmpty
	}
}

func (wildcard KafkaWildcard) ReverseKey4(key string, index1, index2, index3, index4 int) (bool, string, string, string, string){
	if wildcard.Match(key) {
		key := strings.Split(strings.Trim(key, "/"), "/")
		if l:=len(key); l <= index1 || l <= index2 || l <= index3 || l <= index4{
			return false, flagUtils.StrEmpty, flagUtils.StrEmpty, flagUtils.StrEmpty, flagUtils.StrEmpty
		}
		return true, key[index1], key[index2], key[index3], key[index4]
	} else {
		return false, flagUtils.StrEmpty, flagUtils.StrEmpty, flagUtils.StrEmpty, flagUtils.StrEmpty
	}
}
