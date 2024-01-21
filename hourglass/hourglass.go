package hourglass

import (
	"errors"
	"time"
)

// The Hourglass is the tool we use to measure time lapse
type Hourglass struct {
	lastStartTime time.Time
	remainingTime time.Duration
	isRunning     bool
}

// NewHourglass creates an hourglass of a given capacity (duration)
func NewHourglass(d time.Duration) *Hourglass {
	return &Hourglass{
		remainingTime: d,
	}
}

// Start command sets the moment when the hourglass (re)starts
func (h *Hourglass) Start() error {
	if h.isRunning {
		return errors.New("cannot start an hourglass that is already running")
	}
	h.isRunning = true
	h.lastStartTime = time.Now()
	return nil
}

// Stop command stops the hourglass
func (h *Hourglass) Stop() error {
	if !h.isRunning {
		return errors.New("cannot stop an hourglass that is already stopped")
	}
	h.isRunning = false
	h.remainingTime -= h.elapsedTime()
	return nil
}

func (h *Hourglass) elapsedTime() time.Duration {
	return time.Since(h.lastStartTime)
}

// RemainingTime returns the remaining time
func (h *Hourglass) RemainingTime() time.Duration {
	if h.remainingTime-h.elapsedTime() > 0 {
		return (h.remainingTime - h.elapsedTime()).Truncate(time.Second)
	}
	return (0 * time.Second).Truncate(time.Second)
}

// Expired returns whether the hourglass time capacity was totally consumed
func (h *Hourglass) Expired() bool {
	return h.remainingTime <= 0
}
