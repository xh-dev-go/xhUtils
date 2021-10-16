package producer

import "github.com/segmentio/kafka-go"

func Producer(server string) *kafka.Writer {
	w := &kafka.Writer{
		Addr: kafka.TCP(server),
		// NOTE: When Topic is not defined here, each Message must define it instead.
		Balancer: &kafka.LeastBytes{},
	}
	return w
}
