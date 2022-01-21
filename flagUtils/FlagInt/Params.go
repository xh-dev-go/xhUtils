package FlagString

import (
	"flag"
)

type IntParam struct {
	value *int
	name string
	defaultValue int
	usage string
}

func (param *IntParam) Value() int{
	return *param.value
}

func (param *IntParam) Bind(flag *flag.FlagSet) *IntParam {
	flag.IntVar(param.value, param.name, param.defaultValue, param.usage)
	return param
}

func New(name, usage string) *IntParam {
	defaultValue := 0
	v := IntParam{value: &defaultValue}
	v.name = name
	v.usage = usage
	return &v
}

func NewDefault(name, usage string, defaultValue int) *IntParam {
	v := New(name, usage)
	v.defaultValue = defaultValue
	return v
}
