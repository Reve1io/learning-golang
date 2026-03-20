package nil

import "fmt"

type Config struct {
	Host *string
	Port int
}

func PrintConfig(cfg *Config) {

	if cfg.Host == nil {
		fmt.Printf("Host is not set\nPort: %v\n", cfg.Port)
	}

	if cfg.Host != nil {
		fmt.Printf("Host: %v\nPort: %v\n", *cfg.Host, cfg.Port)
	}
}
