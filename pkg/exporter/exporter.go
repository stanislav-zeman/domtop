package exporter

import (
	"context"
	"fmt"
	"log/slog"
	"os"
)

type Exporter struct {
	file           *os.File
	statisticsChan <-chan Serializable
}

func New(file *os.File, dataChan <-chan Serializable) Exporter {
	exporter := Exporter{
		file:           file,
		statisticsChan: dataChan,
	}

	return exporter
}

func (e Exporter) Run(ctx context.Context) {
	for {
		select {
		case statistics, ok := <-e.statisticsChan:
			if !ok {
				e.statisticsChan = nil
			}

			data, err := statistics.Serialize()
			if err != nil {
				slog.Error("could not serialize statics: %v", err)
				continue
			}

			fmt.Fprintf(os.Stdout, data)

		case <-ctx.Done():
		}
	}
}
