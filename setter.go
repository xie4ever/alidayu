package alidayu

// Setter ...
type Setter func(config *config) error

// WithAppKey ...
func WithAppKey(appKey string) Setter {
	return func(config *config) error {
		config.appKey = appKey
		return nil
	}
}

// WithAppSecret ...
func WithAppSecret(appSecret string) Setter {
	return func(config *config) error {
		config.appSecret = appSecret
		return nil
	}
}
