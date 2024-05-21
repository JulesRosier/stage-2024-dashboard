package scheduler

import (
	"log/slog"
	"sync"
	"time"
)

type Scheduler struct {
	wg     *sync.WaitGroup
	doneCh chan bool
}

func NewScheduler() Scheduler {
	var wg sync.WaitGroup
	done := make(chan bool)
	return Scheduler{
		wg:     &wg,
		doneCh: done,
	}
}

func (s Scheduler) Stop() {
	slog.Info("Stopping scheduler")
	close(s.doneCh)
	s.wg.Wait()
}

func (s Scheduler) Schedule(t time.Duration, f func()) {
	s.wg.Add(1)
	ticker := time.NewTicker(t)
	go f()
	go func() {
		for {
			select {
			case <-ticker.C:
				f()
			case <-s.doneCh:
				ticker.Stop()
				s.wg.Done()
				return
			}
		}
	}()
}
