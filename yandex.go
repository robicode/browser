package browser

import (
	"regexp"
	"strings"
)

type Yandex struct {
	id   string
	name string
	ua   string
}

func newYandex(userAgent string) *Yandex {
	return &Yandex{
		id:   "yandex",
		name: "Yandex",
		ua:   userAgent,
	}
}

func (y *Yandex) ID() string {
	return y.id
}

func (y *Yandex) Name() string {
	return y.name
}

func (y *Yandex) Matches() bool {
	return strings.Contains(y.ua, "YaBrowser")
}

func (y *Yandex) FullVersion() string {
	matches := regexp.MustCompile(`YaBrowser/([\d.]+)`).FindStringSubmatch(y.ua)
	if len(matches) > 0 {
		return matches[1]
	}

	return "0.0"
}
