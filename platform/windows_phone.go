package platform

import (
	"regexp"
	"strings"
)

type WindowsPhone struct {
	id        string
	name      string
	userAgent string
}

func NewWindowsPhone(userAgent string) *WindowsPhone {
	return &WindowsPhone{
		id:        "windows_phone",
		name:      "Windows Phone",
		userAgent: userAgent,
	}
}

func (w *WindowsPhone) ID() string {
	return w.id
}

func (w *WindowsPhone) Name() string {
	return w.name
}

func (w *WindowsPhone) Matches() bool {
	return strings.Contains(w.userAgent, "Windows Phone")
}

func (w *WindowsPhone) Version() string {
	return regexp.MustCompile(`Windows Phone ([\d.]+)`).FindString(w.userAgent)
}
