package channel

import "github.com/segmentio/kafka-go"

func GetMessageChan() chan kafka.Message{
	return make(chan kafka.Message)
}
func GetErrChan() chan error{
	return make(chan error)
}

