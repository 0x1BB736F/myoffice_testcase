package main

import (
	"flag"
	"log/slog"
	"os"
	"time"

	"github.com/0x1BB736F/myoffice_testcase/internal/app"
	"github.com/0x1BB736F/myoffice_testcase/pkg/filereader"
)

func main() {
	filePath := flag.String("filepath", "./testdata.txt", "path to file with URLs")
	timeOut := flag.Duration("timeout", time.Second*5, "timeout for every HTTP request")
	flag.Parse()

	if filePath == nil {
		slog.Error("filepath is nil")
		os.Exit(1)
	}

	if timeOut == nil {
		timeout := time.Second * 5
		timeOut = &timeout
	}

	f, err := os.Open(*filePath)
	if err != nil {
		slog.Error(
			"filepath is nil",
			slog.Any("error", err.Error()),
		)
		os.Exit(1)
	}

	cliApp := app.New(
		filereader.New(f),
		app.Config{
			ClientTimeout: *timeOut,
		},
	)

	cliApp.Run()
}
