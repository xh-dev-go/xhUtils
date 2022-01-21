package StringOptions

import (
	"errors"
	"flag"
	"github.com/xh-dev-go/xhUtils/common"
	"github.com/xh-dev-go/xhUtils/flagUtils/flagString"
	"github.com/xh-dev-go/xhUtils/logical"
)

type FlagOptionStringCmd struct {
	list *[]flagString.StringParam
}
type FlagOptionString struct {
	list *[]common.StringTuple
}

func (option *FlagOptionString) Add(key, value string) *FlagOptionString {

	op := common.StringTuple{
		Key: key, Value: value,
	}
	newList := append(
		*option.list,
		op,
	)
	option.list = &newList
	return option
}

func (option *FlagOptionString) Bind(flag *flag.FlagSet) *FlagOptionStringCmd {
	var list []flagString.StringParam
	for _, item := range *option.list {
		list = append(list, *flagString.New(item.Key, item.Value).Bind(flag))
	}
	return &FlagOptionStringCmd{list: &list}
}

var ExDuplicateSelection = errors.New("Duplicate selection")
var ExEmptySelection = errors.New("No selection")
func (option *FlagOptionStringCmd) Value() (string,error) {
	var err error
	var result string
	var bs []bool
	var ti = -1
	for i, item := range *option.list{
		t := "" != item.Value()
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
		l := *option.list
		return (l[ti]).Value(), nil
	}
}

func New() *FlagOptionString {
	return &FlagOptionString{}
}
