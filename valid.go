package valid

import (
	"gopkg.in/validator.v2"
)

func init() {
	validator.SetTag("valid")
}

var errorMap = map[string]string{
	validator.ErrZeroValue.Error():    "字段为空",
	validator.ErrMin.Error():          "小于最小值",
	validator.ErrMax.Error():          "大于最大值",
	validator.ErrLen.Error():          "字段长度错误",
	validator.ErrRegexp.Error():       "字段格式错误",
	validator.ErrUnsupported.Error():  "不支持的字段类型",
	validator.ErrBadParameter.Error(): "tag参数错误",
	validator.ErrUnknownTag.Error():   "tag无法识别",
	validator.ErrInvalid.Error():      "字段值不合法",
}

func Validate(v interface{}) (message string, valid bool) {
	err := validator.Validate(v)
	if err != nil {
		errs := err.(validator.ErrorMap)
		for k, vs := range errs {
			for _, v := range vs {
				message = k + ":" + errorMap[v.Error()]
			}
		}
		return
	}
	valid = true
	return
}
