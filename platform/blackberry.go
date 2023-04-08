package platform

import "regexp"

type BlackBerry struct {
	id        string
	name      string
	userAgent string
}

func NewBlackBerry(userAgent string) *BlackBerry {
	return &BlackBerry{
		id:        "blackberry",
		name:      "BlackBerry",
		userAgent: userAgent,
	}
}

func (bb *BlackBerry) ID() string {
	return bb.id
}

func (bb *BlackBerry) Name() string {
	return bb.name
}

func (bb *BlackBerry) Matches() bool {
	return regexp.MustCompile(`/BB10|BlackBerry`).MatchString(bb.userAgent)
}

func (bb *BlackBerry) Version() string {
	re := regexp.MustCompile(`(?:Version|BlackBerry[\da-z]+)/([\d.]+)`)

	matches := re.FindAllStringSubmatch(bb.userAgent, 1)
	if len(matches) > 0 {
		return matches[0][1]
	}

	return ""
}
