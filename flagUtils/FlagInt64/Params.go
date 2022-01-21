package FlagString

import (
	"flag"
)

type Int64Param struct {
	value *int64
	name string
	defaultValue int64
	usage string
}

func (param *Int64Param) Value() int64{
	return *param.value
}

func (param *Int64Param) Bind(flag *flag.FlagSet) *Int64Param {
	flag.Int64Var(param.value, param.name, param.defaultValue, param.usage)
	return param
}

func New(name, usage string) *Int64Param {
	var defaultValue int64
	defaultValue = 0
	v := Int64Param{value: &defaultValue}
	v.name = name
	v.usage = usage
	return &v
}

func NewDefault(name, usage string, defaultValue int64) *Int64Param {
	v := New(name, usage)
	v.defaultValue = defaultValue
	return v
}
