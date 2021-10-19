package wildcard

import "testing"

func TestKafkaWildcard1(t *testing.T) {
	wildcard := New("a/b/c/d")
	for index, level := range wildcard.levels {
		if index == 0 && level != "a" {
			t.Fail()
		}
		if index == 1 && level != "b" {
			t.Fail()
		}
		if index == 2 && level != "c" {
			t.Fail()
		}
		if index == 3 && level != "d" {
			t.Fail()
		}
	}

}

func TestKafkaWildcard2(t *testing.T) {
	var wildcard = New("a/b/c/d")
	if !wildcard.match("a/b/c/d") {
		t.Fail()
	}
	wildcard = New("a/b/c/*")
	if !wildcard.match("a/b/c/d") {
		t.Fail()
	}
	if wildcard.match("a/b/c/d/e") {
		t.Fail()
	}
	if wildcard.match("a/b/c/") {
		t.Fail()
	}
	if wildcard.match("a/b") {
		t.Fail()
	}
	wildcard = New("a/b/c/>")
	if wildcard.match("a/b") {
		t.Fail()
	}
	if wildcard.match("a/b/c") {
		t.Fail()
	}
	if !wildcard.match("a/b/c/d") {
		t.Fail()
	}
	if !wildcard.match("a/b/c/d/cs") {
		t.Fail()
	}

}
