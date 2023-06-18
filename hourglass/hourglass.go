package hourglass

import "time"

// The Hourglass is the tool we use to measure time lapse
type Hourglass struct {
	creationTime time.Time
	capacity     time.Duration
	running      bool
}

// NewHourglass creates an hourglass of a given capacity (duration)
func NewHourglass(d time.Duration) *Hourglass {
	return &Hourglass{
		capacity: d,
	}
}

// Start command sets the moment when the hourglass starts to measure time lapse
func (h *Hourglass) Start() {
	// protect with mutex?
	h.creationTime = time.Now()
	h.running = true
}

// Stop command stops the hourglass
func (h *Hourglass) Stop() {
	h.running = false
}

func (h *Hourglass) isRunning() bool {
	return h.running
}

func (h *Hourglass) elapsedTime() time.Duration {
	if h.isRunning() {
		return time.Since(h.creationTime)
	}
	return 0 * time.Second
}

// Countdown returns the remaining time
func (h *Hourglass) Countdown() time.Duration {
	if h.isRunning() && h.elapsedTime() < h.capacity {
		return h.capacity - h.elapsedTime() + time.Second
	}
	return 0 * time.Second
}

// Expired returns whether the hourglass time capacity was totally consumed or not
func (h *Hourglass) Expired() bool {
	if h.isRunning() {
		return h.elapsedTime() >= h.capacity
	}
	return false
}
