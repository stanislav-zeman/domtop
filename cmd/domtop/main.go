package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/stanislav-zeman/domtop/pkg/domtop"
	"github.com/stanislav-zeman/domtop/pkg/exporter"
)

var period = flag.String("time", "1s", "domtop refresh period")

func main() {
	config, err := parseArgs()
	if err != nil {
		slog.Error("could not parse command line arguments", "error", err)
		os.Exit(1)
	}

	statisticsChan := make(chan exporter.Serializable, 32)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	exporter := exporter.New(os.Stdout, statisticsChan)
	exporter.Run(ctx)
	domtop, err := domtop.New(config, statisticsChan)
	if err != nil {
		slog.Error("could not start domtop", "error", err)
		os.Exit(1)
	}

	err = domtop.Run(ctx)
	if err != nil {
		slog.Error("failed running domtop", "error", err)
		os.Exit(1)
	}
}

func parseArgs() (config domtop.Config, err error) {
	if len(os.Args) < 2 {
		err = errors.New("missing domain name argument")
		return
	}

	config.Domain = os.Args[1]
	refreshPeriod, err := time.ParseDuration(*period)
	if err != nil {
		err = fmt.Errorf("could not parse time argument: %v", err)
		return
	}

	config.RefreshPeriod = refreshPeriod
	return
}
