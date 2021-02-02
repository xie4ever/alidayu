# alidayu
alidayu-sdk

## 1.alidayu doc

base on https://help.aliyun.com/product/44282.html?spm=a2c4g.11174283.6.540.72b12c42bQEVrz

## 2.all you need

get your sms param:
* app key
* app secret
* sms sign
* template code

## 3.example

```golang

package alidayu

import (
	"fmt"
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

```
