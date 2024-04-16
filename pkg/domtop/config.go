package domtop

import (
	"time"
)

type Config struct {
	Domain        string
	RefreshPeriod time.Duration
}
