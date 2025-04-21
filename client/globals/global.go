package globals

type Config struct {
	Ip       string `json:"ip"`
	Puerto   int    `json:"puerto"`
	Mensaje  string `json:"mensaje"`
	LogLevel string `json:"log_level"`
}

var ClientConfig *Config
