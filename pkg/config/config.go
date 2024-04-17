package config

import (
	"time"
)

type Config struct {
	HypervisorURI string
	DomainName    string
	RefreshPeriod time.Duration
}
