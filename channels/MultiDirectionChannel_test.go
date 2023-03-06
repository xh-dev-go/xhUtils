package channels

import (
	"sync"
	"testing"
)

func TestNewMChannel(t *testing.T) {
	mChan := NewMChannel[string]()

	var result []string

	var wg sync.WaitGroup
	wg.Add(1)

	var blockWg sync.WaitGroup
	blockWg.Add(1)
	go func() {
		blockWg.Done()
		o := mChan.outChan()
		select {
		case x := <-o:
			result = append(result, x)
		}
		select {
		case x := <-o:
			result = append(result, x)
		}
		wg.Done()
	}()
	blockWg.Wait()

	in := mChan.inChan()
	in <- "Hello, world!"
	in <- "Hello, world!2"

	wg.Wait()

	if len(result) != 2 {
		t.Errorf("Expected 2, got %d", len(result))
	}
	if result[0] != "Hello, world!" {
		t.Errorf("Expected 'Hello, world!', got %s", result[0])
	}
	if result[1] != "Hello, world!2" {
		t.Errorf("Expected 'Hello, world!2', got %s", result[1])
	}
}
