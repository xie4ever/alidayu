package alidayu

import (
	"encoding/json"
	"time"
)

// Message 短信本体
type Message struct {
	Method     string `json:"method"`
	AppKey     string `json:"app_key"`
	Timestamp  string `json:"timestamp"`
	Format     string `json:"format"`
	Version    string `json:"v"`
	SignMethod string `json:"sign_method"`
	Type       string `json:"sms_type"`

	FreeSignName string      `json:"sms_free_sign_name"`
	Param        string      `json:"sms_param"`
	Tel          string      `json:"rec_num"`
	TemplateCode string      `json:"sms_template_code"`
	content      interface{} `json:"-"`
}

// NewMessage 新建短信
// signName:  模板名称，可在短信服务获取
func NewMessage(signName string) *Message {
	return &Message{
		Method:       messageMethodSms,
		AppKey:       appKey,
		Timestamp:    time.Now().Format("2006-01-02 15:04:05"),
		Format:       messageFormatJson,
		Version:      messageVersion,
		SignMethod:   messageSignMethodMd5,
		Type:         messageType,
		FreeSignName: signName,
	}
}

// GetTel 获取短信接收方的电话号码
func (m *Message) GetTel() string {
	return m.Tel
}

// SetTel 设置短信接收方的电话号码，暂不支持批量发送
// tel:   接收方电话号码
func (m *Message) SetTel(tel string) *Message {
	m.Tel = tel
	return m
}

// SetTemplateCode 设置短信模板编号
// code:           短信模板编号，可在短信服务获取
func (m *Message) SetTemplateCode(code string) *Message {
	m.TemplateCode = code
	return m
}

// SetContent 设置短信发送信息模板
// content:   短信内容，支持任意对象，但是必须可json序列化，否则发送时报错
func (m *Message) SetContent(content interface{}) (*Message, error) {
	data, err := json.Marshal(content)
	if err != nil {
		return nil, err
	}
	m.content = content
	m.Param = string(data)

	return m, nil
}
