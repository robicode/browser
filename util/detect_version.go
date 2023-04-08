package util

import (
	"reflect"
	"strconv"
	"strings"

	"github.com/robicode/version"
	"github.com/robicode/version/requirement"
)

// DetectVersion checks expected version against actual version
// and returns true if expected satisfies actual.
func DetectVersion(actual, expected any) bool {
	a := parseVersion(actual)
	e := parseVersion(expected)

	if strings.TrimSpace(e) == "" || strings.TrimSpace(a) == "" {
		return false
	}

	r, err := requirement.New(e)
	if err != nil {
		return false
	}

	v, err := version.New(a)
	if err != nil {
		return false
	}

	return r.IsSatisfiedBy(v)
}

func parseVersion(version any) string {
	value := reflect.ValueOf(version)

	if value.Kind() == reflect.String {
		return value.String()
	} else if value.Kind() == reflect.Int {
		return strconv.Itoa(int(value.Int()))
	}

	return ""
}
