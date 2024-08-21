package app

import (
	"context"
	"gis-crawler/pkg/logging"
	"time"
)

type Daemon struct {
	ctx        context.Context
	ticker     *time.Ticker
	interval   time.Duration
	lastHandle time.Time
}

func (d *Daemon) Run(handler func()) {
	go func() {
		for {
			select {
			case <-d.ctx.Done():
				return
			case <-d.ticker.C:
				diff := time.Now().Sub(d.lastHandle)
				logging.Get().Debugf("tick... %v", diff)
				if diff.Seconds() > d.interval.Seconds() {
					//run handler
					handler()

					//set updated time
					d.lastHandle = time.Now()
				}
			}
		}
	}()
}
