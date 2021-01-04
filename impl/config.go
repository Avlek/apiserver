package impl

type Config struct {
	BindAddr string `toml:"bind_addr"`
	LogLevel string `toml:"log_level"`
}

const (
	defaultAddr     = "8087"
	defaultLogLevel = "debug"
)

func NewConfig() *Config {
	return &Config{
		BindAddr: defaultAddr,
		LogLevel: defaultLogLevel,
	}
}
