package valid

import (
	"fmt"
	"testing"
)

type S struct {
	ZeroValue                 string  `valid:"nonzero"`
	MinValue                  int     `valid:"min=0"`
	MaxValue                  int     `valid:"max=1"`
	InvalidLength             string  `valid:"len=0"`
	RegularExpressionMismatch string  `valid:"regexp=^[a-zA-Z]$"`
	unsupportedType           support `valid:"nonzero"`
	BadParameter              string  `valid:"min=a"`
	UnknownTag                string  `valid:"unknowntag=a"`
}

type support struct{}

func Test(t *testing.T) {
	var s S
	s.ZeroValue = "11"
	s.RegularExpressionMismatch = "a"
	message, valid := Validate(s)
	if !valid {
		fmt.Println(message)
	}
	return
}
