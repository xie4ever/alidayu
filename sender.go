package alidayu

import (
	"time"

	"github.com/xiaojiaoyu100/cast"
)

type config struct {
	appKey    string     // key
	appSecret string     // secret
	cast      *cast.Cast // cast
}

// Sender ...
type Sender struct {
	config
}

// NewSender 初始化发送者
// key: alidayu账号分配给应用的AppKey
// secret: alidayu账号分配给应用的AppSecret
func NewSender(setters ...Setter) (*Sender, error) {
	var err error
	s := &Sender{}
	for _, setter := range setters {
		if err := setter(&s.config); err != nil {
			return nil, err
		}
	}
	if err != nil {
		return nil, err
	}

	s.cast, err = cast.New(
		cast.WithBaseURL(productionURL),
		cast.AddCircuitConfig(defaultCircuitName),
		cast.WithDefaultCircuit(defaultCircuitName),
		cast.WithRetry(3),
		cast.WithExponentialBackoffDecorrelatedJitterStrategy(
			time.Millisecond*200,
			time.Millisecond*500,
		),
	)
	return s, err
}
