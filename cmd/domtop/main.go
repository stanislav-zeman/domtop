package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/stanislav-zeman/domtop/pkg/domtop"
)

const defaultRefreshTime = time.Second

func main() {
	if len(os.Args) < 2 {
		log.Fatal("could not start domtop: missing domain name argument")
	}

	refreshTime := defaultRefreshTime
	if len(os.Args) > 2 {
		var err error
		refreshTime, err = time.ParseDuration(os.Args[2])
		if err != nil {
			log.Fatalf("could not parse time argument: %v", err)
		}
	}

	domain := os.Args[1]
	dt := domtop.NewDomtop(domain, refreshTime)
	ctx := context.Background()
	err := dt.Run(ctx)
	if err != nil {
		log.Fatalf("failed running domtop: %v", err)
	}
}
