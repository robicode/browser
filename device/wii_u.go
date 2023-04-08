package device

import "strings"

type WiiU struct {
	id        string
	name      string
	userAgent string
}

func NewWiiU(userAgent string) *WiiU {
	return &WiiU{
		id:        "wiiu",
		name:      "Nintendo WiiU",
		userAgent: userAgent,
	}
}

func (w *WiiU) ID() string {
	return w.id
}

func (w *WiiU) Name() string {
	return w.name
}

func (w *WiiU) Matches() bool {
	return strings.Contains(w.userAgent, "Nintendo WiiU")
}
