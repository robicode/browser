package browser

import (
	"regexp"
	"strings"
)

type QQ struct {
	id   string
	name string
	ua   string
}

func newQQ(userAgent string) *QQ {
	return &QQ{
		id:   "qq",
		name: "QQ Browser",
		ua:   userAgent,
	}
}

func (q *QQ) ID() string {
	return q.id
}

func (q *QQ) Name() string {
	return q.name
}

func (q *QQ) Matches() bool {
	return strings.Contains(q.ua, "QQ") || strings.Contains(q.ua, "QQBrowser")
}

func (q *QQ) FullVersion() string {
	matches := regexp.MustCompile(strings.ToLower(`(?:Mobile MQQBrowser)/([\d.]+)`)).FindStringSubmatch(strings.ToLower(q.ua))
	if len(matches) > 0 {
		return matches[1]
	}

	matches = regexp.MustCompile(strings.ToLower(`(?:QQBrowserLite)/([\d.]+)`)).FindStringSubmatch(strings.ToLower(q.ua))
	if len(matches) > 0 {
		return matches[1]
	}

	matches = regexp.MustCompile(strings.ToLower(`(?:QQBrowser)/([\d.]+)`)).FindStringSubmatch(strings.ToLower(q.ua))
	if len(matches) > 0 {
		return matches[1]
	}

	matches = regexp.MustCompile(strings.ToLower(`(?:QQ)/([\d.]+)`)).FindStringSubmatch(strings.ToLower(q.ua))
	if len(matches) > 0 {
		return matches[1]
	}

	return "0.0"
}
