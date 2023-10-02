package config

var (
	Session struct {
		SecretKey string
		Expire    int
		Issuer    string
	}
)
