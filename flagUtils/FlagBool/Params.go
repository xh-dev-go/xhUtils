package FlagBool

import (
	"flag"
)

type BoolParam struct {
	value *bool
	name string
	defaultValue bool
	usage string
}

func (param *BoolParam) Value() bool{
	return *param.value
}

func (param *BoolParam) Bind(flag *flag.FlagSet) *BoolParam {
	flag.BoolVar(param.value, param.name, param.defaultValue, param.usage)
	return param
}

func New(name, usage string) *BoolParam {
	defaultValue := false
	v := BoolParam{value: &defaultValue}
	v.name = name
	v.usage = usage
	return &v
}

func NewDefault(name, usage string, defaultValue bool) *BoolParam {
	v := New(name, usage)
	v.defaultValue = defaultValue
	return v
}
