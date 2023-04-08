package browser

import (
	"regexp"
	"strings"
)

type Weibo struct {
	id   string
	name string
	ua   string
}

func newWeibo(userAgent string) *Weibo {
	return &Weibo{
		id:   "weibo",
		name: "Weibo",
	}
}

func (w *Weibo) ID() string {
	return w.id
}

func (w *Weibo) Name() string {
	return w.name
}

func (w *Weibo) Matches() bool {
	return strings.Contains(strings.ToLower(w.ua), strings.ToLower("__weibo__"))
}

func (w *Weibo) FullVersion() string {
	matches := regexp.MustCompile(strings.ToLower(`(?:__weibo__)([\d.]+)`)).FindStringSubmatch(strings.ToLower(w.ua))
	if len(matches) > 0 {
		return matches[1]
	}

	return "0.0"
}
