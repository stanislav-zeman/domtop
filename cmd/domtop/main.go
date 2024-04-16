package main

import (
	"context"
	"flag"
	"log"
	"os"
	"time"

	"github.com/stanislav-zeman/domtop/pkg/domtop"
)

var period = flag.String("time", "1s", "domtop refresh period")

func main() {
	config := parseArgs()
	domtop := domtop.NewDomtop(config)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	err := domtop.Run(ctx)
	if err != nil {
		log.Fatalf("failed running domtop: %v", err)
	}
}

func parseArgs() domtop.Config {
	config := domtop.Config{}
	if len(os.Args) < 2 {
		log.Fatal("missing domain name argument")
	}

	config.Domain = os.Args[1]
	refreshPeriod, err := time.ParseDuration(*period)
	if err != nil {
		log.Fatalf("could not parse time argument: %v", err)
	}

	config.RefreshPeriod = refreshPeriod
	return config
}
