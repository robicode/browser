package platform

import (
	"fmt"
	"regexp"
)

type IOS struct {
	id        string
	name      string
	userAgent string
}

const (
	matcher string = `/(iPhone|iPad|iPod)/`

	// The original version matcher from the Ruby gem is incompatible with Go:
	// Original regexp: /OS (?<major>\d+)_(?<minor>\d+)_?(?<patch>\d+)?/
	versionMatcher string = `/OS (\d+)_(\d+)_?(\d+)?/`
)

func NewIOS(userAgent string) *IOS {
	return &IOS{
		id:        "ios",
		userAgent: userAgent,
	}
}

func (i *IOS) ID() string {
	return i.name
}

func (i *IOS) Name() string {
	// TODO finish
	return "iOS"
}

func (i *IOS) Matches() bool {
	return regexp.MustCompile(matcher).MatchString(i.userAgent)
}

func (i *IOS) Version() string {
	re := regexp.MustCompile(versionMatcher)

	matches := re.FindAllStringSubmatch(i.userAgent, -1)

	if len(matches) > 0 {
		fmt.Println(matches)
	}

	// TODO finish

	return ""
}
