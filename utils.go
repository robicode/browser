// This file contains some utilities used by the matchers.
package browser

import (
	"strings"
)

// containsAny returns true if any of the values are found in s.
func containsAny(s string, values ...string) bool {
	for _, v := range values {
		if strings.Contains(s, v) {
			return true
		}
	}
	return false
}
