# valid

```
package valid

import (
	"fmt"
	"testing"
)

type S struct {
	ZeroValue                 string `valid:"nonzero"`           //字段为空
	MinValue                  int    `valid:"min=0"`             //小于最小值
	MaxValue                  int    `valid:"max=1"`             //大于最大值
	InvalidLength             string `valid:"len=0"`             //字段长度错误
	RegularExpressionMismatch string `valid:"regexp=^[a-zA-Z]$"` //字段格式错误
	BadParameter              string `valid:"min=a"`             //tag参数错误
	UnknownTag                string `valid:"unknowntag=a"`      //tag无法识别
}

func Test(t *testing.T) {
	var s S
	s.RegularExpressionMismatch = "a"
	message, valid := Validate(s)
	if !valid {
		fmt.Println(message)
	}
	return
}

输出结果

字段名+错误信息

ZeroValue:字段为空
MinValue:小于最小值
MaxValue:大于最大值
nvalidLength:字段长度错误
RegularExpressionMismatch:字段格式错误
BadParameter:tag参数错误
nknownTag:tag无法识别


```