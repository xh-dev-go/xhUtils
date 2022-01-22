package StringOptions

import (
	"errors"
	"flag"
	"github.com/xh-dev-go/xhUtils/common"
	"github.com/xh-dev-go/xhUtils/flagUtils/FlagSets"
	"github.com/xh-dev-go/xhUtils/flagUtils/flagBool"
	"github.com/xh-dev-go/xhUtils/logical"
)

type FlagOptionBoolCmd struct {
	list []flagBool.BoolParam
}
type FlagOptionBool struct {
	list []common.StringTuple
}

func (option *FlagOptionBool) Add(key, value string) *FlagOptionBool {

	op := common.StringTuple{
		Key: key, Value: value,
	}

	l := option.list
	newList := append(
		l,
		op,
	)
	option.list = newList
	return option
}

func (option *FlagOptionBool) BindCmd() *FlagOptionBoolCmd {
	return option.Bind(FlagSets.CommandFlag)
}

func (option *FlagOptionBool) Bind(flag *flag.FlagSet) *FlagOptionBoolCmd {
	var list []flagBool.BoolParam
	for _, item := range option.list {
		list = append(list, *flagBool.New(item.Key, item.Value).Bind(flag))
	}
	return &FlagOptionBoolCmd{list: list}
}

var ExDuplicateSelection = errors.New("Duplicate selection")
var ExEmptySelection = errors.New("No selection")
func (option *FlagOptionBoolCmd) Value() (string,error) {
	var result string
	var bs []bool
	var ti = -1
	for i, item := range option.list{
		t := item.Value()
		if t {
			if ti == -1 {
				ti = i
			} else {
				return result,ExDuplicateSelection
			}
		}
		bs = append(bs, t)
	}

	logical.OnlyOneOf(bs...)

	if ti == -1 {
		return result, ExEmptySelection
	} else {
		l := option.list
		return (l[ti]).Name(), nil
	}
}

func New() *FlagOptionBool {
	return &FlagOptionBool{}
}
