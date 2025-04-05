package time_frequency

import (
	"sync"
	"time"
)

type TimeFrequency struct {
	interval time.Duration
	lastExec time.Time
	mu       sync.Mutex
}

func NewTimeFrequency(interval time.Duration) *TimeFrequency {
	return &TimeFrequency{
		interval: interval,
		lastExec: time.Now(),
	}
}

func (tfc *TimeFrequency) TryExecute(action func(), fallback func()) {
	tfc.mu.Lock()
	defer tfc.mu.Unlock()

	if time.Since(tfc.lastExec) >= tfc.interval {
		action()
		tfc.lastExec = time.Now()
	} else {
		if fallback != nil {
			fallback()
		}
	}
}
