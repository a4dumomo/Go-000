package pkg

import (
	"errors"
	"log"
)

type ErrCode struct {
	err  error
	code int
}

var CodeInfo = make(map[int]string)

//根据错误码获取错误
func GetErrInfo(code int) string {
	if msg, ok := CodeInfo[code]; ok {
		return msg
	}
	return ""
}

//获取所有错误码
func GetAllCode() []int {
	codes := []int{}
	for k, _ := range CodeInfo {
		codes = append(codes, k)
	}
	return codes
}

//获取所有错误信息
func GetAllErrMsg() []string {
	messages := []string{}
	for _, msg := range CodeInfo {
		messages = append(messages, msg)
	}
	return messages
}

//创建业务错误码
func NewErrCode(msg string, code ...int) *ErrCode {
	err := &ErrCode{err: errors.New(msg)}
	if len(code) == 0 {
		err.code = 0
	} else {
		err.code = code[0]
	}
	if _, ok := CodeInfo[err.code]; ok {
		log.Fatal("Error Code Repetition")
	}
	CodeInfo[err.code] = err.Error()
	return err
}

func (e *ErrCode) Error() string {
	return e.err.Error()
}

func (e *ErrCode) Code() int {
	return e.code
}

func (e *ErrCode) Unwrap() error {
	return e.err
}
