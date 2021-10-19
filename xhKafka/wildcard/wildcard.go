package wildcard

import (
	"context"
	"github.com/segmentio/kafka-go"
	"github.com/xh-dev-go/xhUtils/xhKafka/channel"
	"strings"
	"time"
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
func (wildcard KafkaWildcard) BindRawWildcard(reader kafka.Reader) (chan kafka.Message, chan error) {
	matchChan := channel.GetMessageChan()
	errChan := channel.GetErrChan()
	wildcard.BindWildcard(reader, matchChan, errChan)
	return matchChan, errChan
}
func (wildcard KafkaWildcard) BindWildcard(reader kafka.Reader, matchChan chan kafka.Message, errChan chan error) {
	go func() {
		reader.SetOffsetAt(context.Background(), time.Now())
		for {
			if m, err := reader.ReadMessage(context.Background()); err != nil {
				errChan <- err
				continue
			} else {
				if wildcard.Match(string(m.Key)) {
					matchChan <- m
				}
			}
		}
	}()
}
