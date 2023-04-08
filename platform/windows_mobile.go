package platform

import "strings"

type WindowsMobile struct {
	id        string
	name      string
	userAgent string
}

func NewWindowsMobile(userAgent string) *WindowsMobile {
	return &WindowsMobile{
		id:        "windows_mobile",
		name:      "Windows Mobile",
		userAgent: userAgent,
	}
}

func (w *WindowsMobile) ID() string {
	return w.id
}

func (w *WindowsMobile) Name() string {
	return w.name
}

func (w *WindowsMobile) Version() string {
	return "0"
}

func (w *WindowsMobile) Matches() bool {
	return strings.Contains(w.userAgent, "Windows CE")
}
