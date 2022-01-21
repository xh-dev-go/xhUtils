package flagString

import (
	"flag"
)

type StringParam struct {
	value *string
	name string
	defaultValue string
	usage string
}

func (param *StringParam) IsEmpty() bool {
	return *param.value == ""
}

func (param *StringParam) SetValue(value *string) *StringParam{
	param.value = value
	return param
}

func (param *StringParam) Value() string{
	return *param.value
}

func (param *StringParam) Bind(flag *flag.FlagSet) *StringParam {
	flag.StringVar(param.value, param.name, param.defaultValue, param.usage)
	return param
}

func New(name, usage string) *StringParam {
	defaultValue := ""
	v := StringParam{value: &defaultValue}
	v.name = name
	v.usage = usage
	return &v
}

func NewDefault(name, defaultValue, usage string) *StringParam {
	v := New(name, usage)
	v.defaultValue = defaultValue
	return v
}
