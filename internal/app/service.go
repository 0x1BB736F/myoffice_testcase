package app

import (
	"context"
	"fmt"
	"log/slog"
	"sync"
	"time"

	"github.com/0x1BB736F/myoffice_testcase/pkg/filereader"
	"github.com/0x1BB736F/myoffice_testcase/pkg/httpclient"
)

type App struct {
	reader *filereader.Reader
	pool   sync.Pool

	wg *sync.WaitGroup
}

type Config struct {
	ClientTimeout time.Duration
}

func New(
	reader *filereader.Reader,
	cfg Config,
) *App {
	return &App{
		reader: reader,
		pool: sync.Pool{
			New: func() any {
				r := httpclient.New(cfg.ClientTimeout)
				return r
			},
		},
		wg: &sync.WaitGroup{},
	}
}

func (app *App) Run() {
	readChan := app.reader.ReadChan()

	for URL := range readChan {
		app.wg.Add(1)
		go app.handleURL(URL)
	}

	app.wg.Wait()

	<-time.Tick(time.Second)
	return
}

func (app *App) handleURL(URL string) {
	httpClient := app.pool.Get().(*httpclient.HttpClient)
	defer app.pool.Put(httpClient)
	defer app.wg.Done()

	err := httpClient.VerifyURL(URL)
	if err != nil {
		slog.Error(
			"cant verify URL",
			slog.Any("error", err.Error()),
			slog.String("url", URL),
		)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	res, err := httpClient.Get(ctx, URL)
	cancel()
	if err != nil {
		slog.Error(
			"cant do GET request",
			slog.Any("error", err.Error()),
		)
		return
	}

	slog.Info(fmt.Sprintf("length: %d, duration: %s", res.ContentLength(), res.HandleTime().String()))
}
