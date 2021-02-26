package test

import (
	"testing"

	"alidayu"

	"github.com/pkg/errors"
)

// TestErrEqual ...
func TestErrEqual(t *testing.T) {
	err := errors.New("isv.MOBILE_NUMBER_ILLEGAL")
	t.Log(err == alidayu.MnsError("isv.MOBILE_NUMBER_ILLEGAL"))
	err2 := errors.New("isv.MOBILE_NUMBER_ILLEGAL")
	t.Log(err == err2)
}

// TestErrEqualV2 ...
func TestErrEqualV2(t *testing.T) {
	err := errors.New("isv.MOBILE_NUMBER_ILLEGAL")
	e, ok := err.(alidayu.MnsError)
	if !ok {
		t.Log(ok)
	}
	t.Log(e == "isv.MOBILE_NUMBER_ILLEGAL")
}

// TestInvalidTel 测试非法电话号码
func TestInvalidTel(t *testing.T) {
	msg, _ := s.NewMessage(testSignName).
		SetTemplateCode(testTemplateCode).
		SetTel("1509996624").
		SetContent(map[string]string{
			"code": "fuck",
			"min":  "5",
		})
	err := s.Send(msg)
	if err != nil {
		if alidayu.IsSlightError(err) {
			return
		}
		t.Fatal(err)
	}
}

// TestEmptySignName 测试空短信签名
func TestEmptySignName(t *testing.T) {
	msg, _ := s.NewMessage("").
		SetTemplateCode(testTemplateCode).
		SetTel(testTel).
		SetContent(map[string]string{
			"customer": "xie4ever",
		})
	err := s.Send(msg)
	if err != nil {
		if alidayu.IsSevereError(err) {
			return
		}
		t.Fatal(err)
	}
}

// TestEmptyTemplateCode 测试空模板号
func TestEmptyTemplateCode(t *testing.T) {
	msg, _ := s.NewMessage(testSignName).
		SetTemplateCode("").
		SetTel(testTel).
		SetContent(map[string]string{
			"customer": "xie4ever",
		})
	err := s.Send(msg)
	if err != nil {
		if alidayu.IsSevereError(err) {
			return
		}
		t.Fatal(err)
	}
}

// TestEmptyTel 测试空电话号码
func TestEmptyTel(t *testing.T) {
	msg, _ := s.NewMessage(testSignName).
		SetTemplateCode(testTemplateCode).
		SetTel("").
		SetContent(map[string]string{
			"customer": "xie4ever",
		})
	err := s.Send(msg)
	if err != nil {
		if alidayu.IsSlightError(err) {
			return
		}
		t.Fatal(err)
	}
}

// TestEmptyParam 测试空参数
func TestEmptyParam(t *testing.T) {
	msg, err := s.NewMessage(testSignName).
		SetTemplateCode(testTemplateCode).
		SetTel(testTel).
		SetContent(nil)
	if err != nil {
		t.Fatal(err)
	}
	err = s.Send(msg)
	if err != nil {
		if alidayu.IsSevereError(err) {
			return
		}
		t.Fatal(err)
	}
}
