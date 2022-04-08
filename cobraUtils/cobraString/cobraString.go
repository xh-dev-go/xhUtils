package cobraString

import "github.com/spf13/cobra"

type CobraString struct {
	value        *string
	name         string
	defaultValue string
	usage        string
	shortHand    string
}

func (param *CobraString) Name() string {
	return param.name
}
func (param *CobraString) SetValue(value *string) *CobraString {
	param.value = value
	return param
}

func (param *CobraString) Value() string {
	return *param.value
}

func (param *CobraString) Shorthand(shorthand string) *CobraString {
	param.shortHand = shorthand
	return param
}

func (param *CobraString) Bind(cmd *cobra.Command) *CobraString {
	if param.shortHand == "" {
		cmd.Flags().StringVar(param.value, param.name, param.defaultValue, param.usage)
	} else {
		cmd.Flags().StringVarP(param.value, param.name, param.shortHand, param.defaultValue, param.usage)
	}
	return param
}

func New(name, usage string) *CobraString {
	defaultValue := ""
	v := CobraString{value: &defaultValue}
	v.name = name
	v.usage = usage
	return &v
}

func NewDefault(name, usage string, defaultValue string) *CobraString {
	v := New(name, usage)
	v.defaultValue = defaultValue
	return v
}
