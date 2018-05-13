package app

type Config struct {
	ListenAddress string `toml:"listen_address`
}

func NewDefaultConfig() *Config {
	return &Config{
		ListenAddress: "127.0.0.1:8080",
	}
}
