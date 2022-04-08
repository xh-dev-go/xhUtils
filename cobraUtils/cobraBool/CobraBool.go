package cobraBool

import (
	"github.com/spf13/cobra"
)

type CobraBool struct {
	value        *bool
	name         string
	defaultValue bool
	usage        string
	shortHand    string
}

func (param *CobraBool) Name() string {
	return param.name
}
func (param *CobraBool) SetValue(value *bool) *CobraBool {
	param.value = value
	return param
}

func (param *CobraBool) Value() bool {
	return *param.value
}

func (param *CobraBool) Shorthand(shorthand string) *CobraBool {
	param.shortHand = shorthand
	return param
}

func (param *CobraBool) Bind(cmd *cobra.Command) *CobraBool {
	if param.shortHand == "" {
		cmd.Flags().BoolVar(param.value, param.name, param.defaultValue, param.usage)
	} else {
		cmd.Flags().BoolVarP(param.value, param.name, param.shortHand, param.defaultValue, param.usage)
	}
	return param
}

func New(name, usage string) *CobraBool {
	defaultValue := false
	v := CobraBool{value: &defaultValue}
	v.name = name
	v.usage = usage
	return &v
}

func NewDefault(name, usage string, defaultValue bool) *CobraBool {
	v := New(name, usage)
	v.defaultValue = defaultValue
	return v
}
