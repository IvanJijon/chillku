package ui

type UI struct {
	Header    *Header
	Countdown *Countdown
}

func NewUI() UI {
	return UI{
		Header:    InitHeader(),
		Countdown: InitCountdown(),
	}
}
