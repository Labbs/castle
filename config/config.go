package config

var (
	ConfigFile     string
	Debug          bool
	Port           int
	Version        string
	EnableHTTPLogs bool

	Database struct {
		DSN    string
		Engine string
	}
)
