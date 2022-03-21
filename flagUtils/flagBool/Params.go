package flagBool

import (
	"flag"
	"github.com/xh-dev-go/xhUtils/flagUtils/FlagSets"
)

type BoolParam struct {
	value *bool
	name string
	defaultValue bool
	usage string
}

func (param *BoolParam) Name() string {
	return param.name
}
func (param *BoolParam) SetValue(value *bool) *BoolParam{
	param.value = value
	return param
}

func (param *BoolParam) Value() bool{
	return *param.value
}

func (param *BoolParam) Bind(flag *flag.FlagSet) *BoolParam {
	flag.BoolVar(param.value, param.name, param.defaultValue, param.usage)
	return param
}
func (param *BoolParam) BindCmd() *BoolParam {
	return param.Bind(FlagSets.CommandFlag)
}

func (param *BoolParam) Share(name, usage string) *BoolParam {
	newParam := BoolParam{
		value: param.value,
		name: name,
		usage: usage,
	}
	return &newParam
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
