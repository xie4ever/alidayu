package alidayu

import (
	"strings"
	"time"

	"github.com/xiaojiaoyu100/cast"
)

// InitAlidayu 初始化服务
// key:        alidayu账号分配给应用的AppKey
// secret:     alidayu账号分配给应用的AppSecret
// cast:       github.com/xiaojiaoyu100/cast中的cast
// blackList:  alidayu禁止发送的短信模板黑名单
func InitAlidayu(key, secret string) error {
	appKey = key
	appSecret = secret

	var err error
	c, err = cast.New(
		cast.WithBaseURL(productionURL),
		cast.AddCircuitConfig(defaultCircuitName),
		cast.WithDefaultCircuit(defaultCircuitName),
		cast.WithRetry(3),
		cast.WithExponentialBackoffDecorrelatedJitterStrategy(
			time.Millisecond*200,
			time.Millisecond*500,
		),
	)
	return err
}

// SendMessage 发送邮件
// message:    短信模板
func SendMessage(message *Message) error {
	// 校验AppKey
	if strings.TrimSpace(appKey) == "" {
		return appKeyIsEmpty
	}
	// 校验AppSecret
	if strings.TrimSpace(appSecret) == "" {
		return appSecretIsEmpty
	}
	// 校验SignName
	if strings.TrimSpace(message.FreeSignName) == "" {
		return freeSignNameIsEmpty
	}
	// 校验Code
	if strings.TrimSpace(message.TemplateCode) == "" {
		return templateCodeIsEmpty
	}
	// 校验Tel
	if strings.TrimSpace(message.Tel) == "" {
		return telIsEmpty
	}
	// 校验cast
	if c == nil {
		return castIsNil
	}

	return post(c, message)
}
