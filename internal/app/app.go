package app

import (
	"context"
	"gis-crawler/internal/config"
	"gis-crawler/internal/rest"
	"gis-crawler/internal/service/gis"
	"gis-crawler/internal/storage"
	"time"
)

type App struct {
	cancel  context.CancelFunc
	cancel2 context.CancelFunc
	d       *Daemon
}

func New(cfg *config.Config) *App {
	ctx, cancel := context.WithCancel(context.Background())
	app := App{
		cancel: cancel,
	}

	d := Daemon{
		ctx:      ctx,
		ticker:   time.NewTicker(time.Second * 3),
		interval: time.Minute * time.Duration(cfg.Interval),
	}
	app.d = &d

	//register storage
	strg := storage.MustLoad(ctx, cfg)

	//register rest client
	gisClient := rest.NewGisClient(ctx, cfg.Gis.Host)

	//Register Service
	srvc := gis.New(strg, gisClient, cfg.Gis)

	//Start handlers
	d.Run(srvc.Handle)

	return &app
}

func (app *App) Stop() {
	app.cancel()
}
