package KKey

import (
	"github.com/xh-dev-go/xhUtils/xhKafka/KWildcard"
	"strings"
)

type KKey struct {
	segments []string
}

func FromKey(key string) *KKey{
	segments := strings.Split(strings.Trim(key, "/"), "/")
	return &KKey{
		segments: segments,
	}
}

func (key *KKey) Push(segment string) *KKey{
	return &KKey{
		segments: append(key.segments, segment),
	}
}

func (key *KKey) Event() *KKey{
	return &KKey{
		segments: append(key.segments, "evt"),
	}
}
func (key *KKey) Request() *KKey{
	return &KKey{
		segments: append(key.segments, "req"),
	}
}
func (key *KKey) Response() *KKey{
	return &KKey{
		segments: append(key.segments, "resp"),
	}
}
func (key *KKey) Cmd() *KKey{
	return &KKey{
		segments: append(key.segments, "cmd"),
	}
}

func (key *KKey) Pop() *KKey {
	return &KKey{
		segments: key.segments[:len(key.segments)-1],
	}
}

func (key *KKey) ToKey() string{
	return "/"+strings.Join(key.segments,"/")
}

func (key *KKey) ToWildcard() KWildcard.KafkaWildcard {
	return KWildcard.New(key.ToKey())
}
