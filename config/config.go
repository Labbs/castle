package config

var AppConfig Config = Config{}

type Config struct {
	ConfigFile     string
	Debug          bool
	PrettyLogs     bool
	LocalDev       bool
	Port           int
	Version        string
	EnableHTTPLogs bool

	InitFrontendViewEngine bool

	Database struct {
		DSN    string
		Engine string
	}

	Session struct {
		SecretKey string
		Expire    int
		Issuer    string
	}
}
