package alidayu

import "strings"

// Send 发送短信
// message: 短信模板
func (s *Sender) Send(m *message) error {
	if err := s.check(m); err != nil {
		return err
	}
	return s.post(m)
}

func (s *Sender) check(message *message) error {
	if strings.TrimSpace(s.appKey) == "" {
		return appKeyIsEmpty
	}
	if strings.TrimSpace(s.appSecret) == "" {
		return appSecretIsEmpty
	}
	if strings.TrimSpace(message.FreeSignName) == "" {
		return freeSignNameIsEmpty
	}
	if strings.TrimSpace(message.TemplateCode) == "" {
		return templateCodeIsEmpty
	}
	if strings.TrimSpace(message.Tel) == "" {
		return telIsEmpty
	}
	if s.cast == nil {
		return castIsNil
	}
	return nil
}
