package config

import (
	"time"
)

type Config struct {
	HypervisorURL string
	DomainName    string
	RefreshPeriod time.Duration
}
