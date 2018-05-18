package app

// Config contains config data for the server
type Config struct {
	ListenAddress string
}

func NewDefaultConfig() *Config {
	return &Config{
		ListenAddress: "127.0.0.1:8080",
	}
}
