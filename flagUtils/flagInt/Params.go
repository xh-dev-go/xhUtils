package flagInt

import (
	"flag"
	"github.com/xh-dev-go/xhUtils/flagUtils/FlagSets"
)

type IntParam struct {
	value *int
	name string
	defaultValue int
	usage string
}

func (param *IntParam) Name() string {
	return param.name
}
func (param *IntParam) SetValue(value *int) *IntParam{
	param.value = value
	return param
}

func (param *IntParam) Value() int{
	return *param.value
}

func (param *IntParam) Bind(flag *flag.FlagSet) *IntParam {
	flag.IntVar(param.value, param.name, param.defaultValue, param.usage)
	return param
}
func (param *IntParam) BindCmd() *IntParam {
	return param.Bind(FlagSets.CommandFlag)
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
