package cobraInt

import (
	"github.com/spf13/cobra"
)

type CobraInt struct {
	value        *int
	name         string
	defaultValue int
	usage        string
	shortHand    string
}

func (param *CobraInt) Name() string {
	return param.name
}
func (param *CobraInt) SetValue(value *int) *CobraInt {
	param.value = value
	return param
}

func (param *CobraInt) Value() int {
	return *param.value
}

func (param *CobraInt) Shorthand(shorthand string) *CobraInt {
	param.shortHand = shorthand
	return param
}

func (param *CobraInt) Bind(cmd *cobra.Command) *CobraInt {
	if param.shortHand == "" {
		cmd.Flags().IntVar(param.value, param.name, param.defaultValue, param.usage)
	} else {
		cmd.Flags().IntVarP(param.value, param.name, param.shortHand, param.defaultValue, param.usage)
	}
	return param
}

func New(name, usage string) *CobraInt {
	defaultValue := 0
	v := CobraInt{value: &defaultValue}
	v.name = name
	v.usage = usage
	return &v
}

func NewDefault(name, usage string, defaultValue int) *CobraInt {
	v := New(name, usage)
	v.defaultValue = defaultValue
	return v
}
