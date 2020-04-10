package config

import "github.com/kelseyhightower/envconfig"

// Config is a strutcture to set the server behavior
type Config struct {
	Port       string `envconfig: "PORT" required:"true"`
	Cookies    bool   `envconfig: "COOKIES"`
	UseWebhook bool   `envconfig: "USE_WEBHOOK" required:"true"`
}

// Setup initialize bot configuration vars
func Setup() (*Config, error) {
	var c Config
	if err := envconfig.Process("", &c); err != nil {
		return &c, err
	}
	return &c, nil
}
