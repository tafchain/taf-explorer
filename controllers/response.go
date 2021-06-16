package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type Error interface {
	Code() int
	Error() string
}

func Err(code int, text string) Error {
	return &errorString{code: code, s: text}
}

type errorString struct {
	code int
	s    string
}

func (e *errorString) Code() int {
	return e.code
}

func (e *errorString) Error() string {
	return e.s
}

type resp struct {
	Code int               `json:"code"`
	Data interface{}       `json:"data,omitempty"`
	Msg  string            `json:"msg"`
	c    *beego.Controller `json:"-"`
}

func (r *resp) Ok(data interface{}) {
	r.Data = data
	r.c.Data["json"] = r
	r.c.ServeJSON()
}

func (r *resp) Err(err error) {
	if err == nil {
		r.c.Abort("err is nil")
		return
	}
	r.Msg = err.Error()
	if e, ok := err.(Error); ok {
		r.Code = e.Code()
	} else {
		r.Code = 1
	}
	logs.Error("resp: ", "code=", r.Code, ",msg=", r.Msg, ",err=", err)
	r.c.Data["json"] = r
	r.c.ServeJSON()
}

func (r *resp) ErrWith(err error, code int, msg string) {
	if err == nil {
		r.Error(code, msg, nil)
		return
	}
	if e, ok := err.(Error); ok {
		r.Error(e.Code(), e.Error(), e)
		return
	}
	r.Error(code, msg, err)
}

func (r *resp) Error(code int, msg string, err error) {
	r.Code = code
	r.Msg = msg
	if err != nil {
		logs.Error("resp: ", "code=", code, ",msg=", r.Msg, ",err=", err)
	}
	r.c.Data["json"] = r
	r.c.ServeJSON()
}

func (r *resp) ErrParamUnmarshal(err error) {
	if err == nil {
		r.Error(paramUnmarshalErr, "param unmarshal failed", nil)
		return
	}
	if e, ok := err.(Error); ok {
		r.Error(e.Code(), e.Error(), e)
		return
	}
	r.Error(paramUnmarshalErr, "param unmarshal failed", err)
}

const (
	paramUnmarshalErr = 100000
	BlockReqParamErr = 100100 + iota
	BlockDbErr
)
