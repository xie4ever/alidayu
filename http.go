package alidayu

import (
	"bytes"
	"context"
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"sort"
	"strings"
	"time"
)

type postResp struct {
	ErrorResponse errorResponse `json:"error_response"`
}

type errorResponse struct {
	Code      int    `json:"code"`
	SubCode   string `json:"sub_code"`
	Msg       string `json:"msg"`
	SubMsg    string `json:"sub_msg"`
	RequestID string `json:"request_id"`
}

func (s *Sender) post(message *message) error {
	body, _, err := s.getRequestBody(message)
	if err != nil {
		return getRequestBodyFailed
	}

	request := s.cast.NewRequest().
		WithCustomBody("application/x-www-form-urlencoded", body).
		WithTimeout(5 * time.Second).
		Post()
	response, err := s.cast.Do(context.TODO(), request)
	if err != nil {
		return castRequestFailed
	}

	// 尝试序列化
	var postResp postResp
	err = json.Unmarshal(response.Body(), &postResp)
	if err == nil && postResp.ErrorResponse.Code != 0 {
		// 如果序列化成功，识别错误码
		return errors.New(postResp.ErrorResponse.SubCode)
	}

	// 如果序列化失败，说明成功
	return nil
}

func (s *Sender) getRequestBody(message *message) ([]byte, int64, error) {
	m, err := structToStringMap(message)
	if err != nil {
		return nil, 0, err
	}

	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	v := url.Values{}
	var sign bytes.Buffer

	_, err = sign.WriteString(s.appSecret)
	if err != nil {
		return nil, 0, err
	}

	for _, k := range keys {
		if len(m[k]) == 0 {
			continue
		}
		v.Set(k, m[k])
		sign.WriteString(k + m[k])
	}

	_, err = sign.WriteString(s.appSecret)
	if err != nil {
		return nil, 0, err
	}

	v.Set("sign", strings.ToUpper(fmt.Sprintf("%x", md5.Sum(sign.Bytes()))))

	return []byte(v.Encode()), int64(len(v.Encode())), nil
}
