package browser

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type InternetExplorer struct {
	id        string
	name      string
	userAgent string
}

var tridentMapping = map[string]string{
	"4.0": "8",
	"5.0": "9",
	"6.0": "10",
	"7.0": "11",
	"8.0": "12",
}

func newInternetExplorer(userAgent string) *InternetExplorer {
	return &InternetExplorer{
		id:        "ie",
		name:      "Internet Explorer",
		userAgent: userAgent,
	}
}

func (i *InternetExplorer) ID() string {
	return i.id
}

func (i *InternetExplorer) Name() string {
	return i.name
}

func (i *InternetExplorer) Matches() bool {
	return isMSIE(i.userAgent) || i.isModernIE()
}

func isMSIE(userAgent string) bool {
	return strings.Contains(userAgent, "MSIE") && !strings.Contains(userAgent, "Opera")
}

func (i *InternetExplorer) isModernIE() bool {
	return regexp.MustCompile(`Trident/.*?; rv:(.*?)`).MatchString(i.userAgent)
}

func (i *InternetExplorer) FullVersion() string {
	return fmt.Sprintf("%s.0", i.ieVersion())
}

func (i *InternetExplorer) MSIEFullVersion() string {
	matches := regexp.MustCompile(`MSIE ([\d.]+)|Trident/.*?; rv:([\d.]+)`).FindStringSubmatch(i.userAgent)
	if len(matches) > 0 {
		if len(matches) > 1 {
			return matches[2]
		} else {
			return matches[1]
		}
	} else {
		return "0.0"
	}
}

func (i *InternetExplorer) MSIEVersion() string {
	return strings.Split(i.MSIEFullVersion(), ".")[0]
}

func (i *InternetExplorer) tridentVersion() string {
	matches := regexp.MustCompile(`Trident/([0-9.]+)`).FindStringSubmatch(i.userAgent)
	if len(matches) > 0 {
		return matches[len(matches)-1]
	} else {
		return "0"
	}
}

func (i *InternetExplorer) ieVersion() string {
	v, ok := tridentMapping[i.tridentVersion()]
	if !ok {
		return i.MSIEVersion()
	} else {
		return v
	}
}

func (i *InternetExplorer) IsCompatibilityView() bool {
	tv, err := strconv.Atoi(i.tridentVersion())
	if err != nil {
		return false
	}

	iv, err := strconv.Atoi(i.MSIEVersion())
	if err != nil {
		return false
	}

	if tv > 0 {
		return iv < (tv + 4)
	}

	return false
}
