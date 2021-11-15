package retcode

import (
	"encoding/json"
	"fmt"
)

type Error struct {
	RetCode string `json:"ret_code"`
	RetDesc string `json:"ret_desc"`
}

var codes = map[string]string{}

func NewError(code, desc string) *Error {
	if v, ok := codes[code]; ok {
		panic(fmt.Sprintf("错误码：%s已存在，对应描述为：%s，请更换！", code, v))
	}
	codes[code] = desc
	return &Error{RetCode: code, RetDesc: desc}
}

func (e Error) Marshal() []byte {
	data, _ := json.Marshal(e)
	return data
}
