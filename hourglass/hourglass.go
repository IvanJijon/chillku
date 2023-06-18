package hourglass

import "time"

type Hourglass struct {
	ticker       *time.Ticker
	creationTime time.Time
	d            time.Duration
	running      bool
}

func NewHourglass(d time.Duration) *Hourglass {
	return &Hourglass{
		d: d,
	}
}

func (h *Hourglass) Start() {
	// protect with mutex?
	h.creationTime = time.Now()
	h.ticker = time.NewTicker(time.Millisecond)
	h.running = true
}

func (h *Hourglass) Stop() {
	h.ticker.Stop()
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

func (h *Hourglass) CountDown() time.Duration {
	if h.isRunning() {
		return h.d - h.elapsedTime()
	}
	return 0 * time.Second
}
func (h *Hourglass) Expired() bool {
	if h.isRunning() {
		return h.elapsedTime() >= h.d
	}
	return false
}
