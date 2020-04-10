package app

import (
	"fmt"

	config "github.com/whoismath/ademir/config"
)

// AdemirBot is the principal bot structure
type AdemirBot struct {
	conf *config.Config
}

// CreateApp initiates AdemirBot
func CreateApp(conf *config.Config) *AdemirBot {
	return &AdemirBot{conf}
}

// Start inializes Application
func (AdemirBot *AdemirBot) Start() {
	fmt.Println("eae")
}
