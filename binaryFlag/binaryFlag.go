package binaryFlag

import (
	"errors"
	"fmt"
	"math"
)

type BinaryFlag struct {
	value int
}

const FlagSize = 32 - 1

func (flag *BinaryFlag) UnSetBit(bit int) *BinaryFlag {
	if bit < 0 {
		panic(errors.New("bit can not be negative"))
	} else if bit < FlagSize {
		flag.value &= ^(1 << bit)
		return flag
	} else {
		panic(errors.New("bit can not bet larger than 64"))
	}
}

func (flag *BinaryFlag) SetBit(bit int) *BinaryFlag {
	if bit < 0 {
		panic(errors.New("bit can not be negative"))
	} else if bit <= FlagSize {
		flag.value |= 1 << bit
		return flag
	} else {
		panic(errors.New(fmt.Sprintf("bit can not bet larger than 64[%d]", bit)))
	}
}

func (flag *BinaryFlag) IsSet(bit int) bool {
	if bit < 0 {
		panic(errors.New("bit can not be negative"))
	} else if bit <= FlagSize {
		v := 1 << bit

		return (flag.value & v) == v
	} else {
		panic(errors.New(fmt.Sprintf("bit can not bet larger than 64[%d]", bit)))
	}
}

func (flag *BinaryFlag) Value() int {
	return flag.value
}

func (flag *BinaryFlag) SetBinary(setFlag BinaryFlag) *BinaryFlag {
	flag.value |= setFlag.value
	return flag
}

func (flag *BinaryFlag) AllMatch(setFlag BinaryFlag) bool {
	return flag.value&setFlag.value == setFlag.value
}

func (flag *BinaryFlag) AnyMatch(setFlag BinaryFlag) bool {
	if setFlag.value > 0 {
		return flag.value&setFlag.value > 0
	} else {
		return false
	}
}

func (flag *BinaryFlag) Toggle() *BinaryFlag {
	flag.value ^= math.MaxUint >> 1
	return flag
}

func (flag *BinaryFlag) DefaultValueStr() string {
	return flag.ValueStr("X", "_")
}
func (flag *BinaryFlag) ValueStr(on, off string) string {
	var s = ""
	for i := 0; i < FlagSize; i++ {
		if flag.IsSet(i) {
			s = on + s
		} else {
			s = off + s
		}
	}
	return s
}

type ValuePair[T any] struct {
	values map[int]T
}

func (pair ValuePair[T]) ExtractAny(flag BinaryFlag) (T, error) {
	var v T
	for i, d := range pair.values {
		if flag.IsSet(i) {
			return d, nil
		}
	}
	return v, errors.New("not found")
}

func (pair ValuePair[T]) ExtractAll(flag BinaryFlag) []T {
	var arr []T
	for i, d := range pair.values {
		if flag.IsSet(i) {
			arr = append(arr, d)
		}
	}
	return arr
}

func New() *BinaryFlag {
	return &BinaryFlag{}
}
