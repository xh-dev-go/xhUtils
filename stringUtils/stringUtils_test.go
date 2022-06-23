package stringUtils

import (
	"testing"
)

func TestStringConcat(t *testing.T) {
	s := StringConcat("sb", " x-x ", "bs", true)
	if s != "sb x-x bs\n" {
		t.Fail()
	}
	s = StringConcat("sb", " x-x ", "bs", false)
	if s != "sb x-x bs" {
		t.Fail()
	}
	s = StringConcat("b", "\t", "s", false)
	if s != "b\ts" {
		t.Fail()
	}
}

func TestRemoveFirst(t *testing.T) {
	s, b := RemoveFirst("sb", "sss")
	if b {
		t.Fail()
	}
	if s != "sb" {
		t.Fail()
	}

	s, b = RemoveFirst("sb", "s")
	if !b {
		t.Fail()
	}
	if s != "b" {
		t.Fail()
	}

	s, b = RemoveFirst("xxxxsb", "s")
	if !b {
		t.Fail()
	}
	if s != "b" {
		t.Fail()
	}

}
