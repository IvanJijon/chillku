package ui

type ChillkuUI struct {
	Header    *Header
	Countdown *Countdown
}

func NewChillkuUI() ChillkuUI {
	return ChillkuUI{
		Header:    InitHeader(),
		Countdown: InitCountdown(),
	}
}
