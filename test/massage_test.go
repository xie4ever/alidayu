package alidayu

import (
	"testing"

	"alidayu"
)

// API错误码请见https://helpcdn.aliyun.com/document_detail/101346.html
const (
	testKey          = "your app key"       // 分配给应用的AppKey
	testSecret       = "yout test secret"   // 分配给应用的AppSecret
	testSignName     = "your sign name"     // 短信签名
	testTemplateCode = "your template code" // 短信模板ID
	testTel          = "your telephone"     // 接收手机号码
)

func init() {
	alidayu.InitAlidayu(testKey, testSecret)
}

// TestNewMessage 测试短信模板
func TestNewMessage(t *testing.T) {
	msg, _ := alidayu.NewMessage(testSignName).
		SetTemplateCode(testTemplateCode).
		SetTel(testTel).
		SetContent(map[string]string{
			"code": "fuck",
			"min":  "5",
		})
	t.Log(msg)
}

// TestSendMessageInMap 测试使用Map发送单条短信
func TestSendMessageInMap(t *testing.T) {
	msg, _ := alidayu.NewMessage(testSignName).
		SetTemplateCode(testTemplateCode).
		SetTel(testTel).
		SetContent(map[string]string{
			"code": "fuck",
			"min":  "5",
		})
	err := alidayu.SendMessage(msg)
	if err != nil {
		t.Fatal(err)
	}
}

// TestSendMessageInObj 测试使用对象发送单条短信
func TestSendMessageInObj(t *testing.T) {
	type Content struct {
		Customer string `json:"customer"`
	}
	msg, _ := alidayu.NewMessage(testSignName).
		SetTemplateCode(testTemplateCode).
		SetTel(testTel).
		SetContent(Content{
			Customer: "xie4ever",
		})
	err := alidayu.SendMessage(msg)
	if err != nil {
		t.Fatal(err)
	}
}
