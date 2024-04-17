package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/spf13/pflag"
	"github.com/stanislav-zeman/domtop/pkg/config"
	"github.com/stanislav-zeman/domtop/pkg/domtop"
	"github.com/stanislav-zeman/domtop/pkg/exporter"
	"github.com/stanislav-zeman/domtop/pkg/statistics"
)

func main() {
	config, err := parseArgs()
	if err != nil {
		slog.Error("could not parse command line arguments", "error", err)
		os.Exit(1)
	}

	statisticsChan := make(chan statistics.Serializable, 32)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	exporter := exporter.New(os.Stdout, statisticsChan)
	go exporter.Run(ctx)
	domtop, err := domtop.New(config, statisticsChan)
	if err != nil {
		slog.Error("could not start domtop", "error", err)
		os.Exit(1)
	}

	domtop.Run(ctx)
}

func parseArgs() (cfg config.Config, err error) {
	period := pflag.String("time", "1s", "domtop refresh period")
	hypervisorURI := pflag.String("hypervisor-uri", "qemu:///system", "libvirt hypervisor uri")
	pflag.Parse()
	if len(os.Args) < 2 {
		err = errors.New("missing domain name argument")
		return
	}

	cfg.DomainName = os.Args[1]
	refreshPeriod, err := time.ParseDuration(*period)
	if err != nil {
		err = fmt.Errorf("could not parse time argument: %v", err)
		return
	}

	cfg.RefreshPeriod = refreshPeriod
	cfg.HypervisorURI = *hypervisorURI
	return
}
