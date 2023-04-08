package browser

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/robicode/browser/platform"
	"github.com/robicode/browser/util"
	"github.com/robicode/version"
)

type Platform struct {
	id        string
	name      string
	version   string
	userAgent string
	subject   PlatformMatcher

	matchers []PlatformMatcher
}

// PlatformMatcher is the interface for a Platform matcher.
type PlatformMatcher interface {
	ID() string
	Matches() bool
	Name() string
	Version() string
}

// NewPlatform creates a newly initialised *Platform
func NewPlatform(userAgent string) *Platform {
	adobeAir := platform.NewAdobeAir(userAgent)
	chromeOS := platform.NewChromeOS(userAgent)
	winMobile := platform.NewWindowsMobile(userAgent)
	winPhone := platform.NewWindowsPhone(userAgent)
	android := platform.NewAndroid(userAgent)
	bb := platform.NewBlackBerry(userAgent)
	ios := platform.NewIOS(userAgent)
	mac := platform.NewMac(userAgent)
	kos := platform.NewKaiOS(userAgent)
	ffos := platform.NewFirefoxOS(userAgent)
	win := platform.NewWindows(userAgent)
	linux := platform.NewLinux(userAgent)
	unknown := platform.NewUnknown(userAgent)

	defaultMatchers := []PlatformMatcher{
		adobeAir,
		chromeOS,
		winMobile,
		winPhone,
		android,
		bb,
		ios,
		mac,
		kos,
		ffos,
		win,
		linux,
		unknown,
	}

	var id string
	var name string
	var version string
	var sub PlatformMatcher

	for _, matcher := range defaultMatchers {
		if matcher.Matches() {
			id = matcher.ID()
			name = matcher.Name()
			version = matcher.Version()
			sub = matcher
			break
		}
	}

	return &Platform{
		id:        id,
		name:      name,
		version:   version,
		userAgent: userAgent,
		matchers:  defaultMatchers,
		subject:   sub,
	}
}

func (p *Platform) IsAdobeAir() bool {
	return p.id == "adobe_air"
}

func (p *Platform) IsAndroid() bool {
	return p.id == "android"
}

func (p *Platform) IsAndroidVersion(major, minor int) bool {
	requested, err := version.New(fmt.Sprintf("%d.%d", major, minor))
	if err != nil {
		return false
	}

	actual, err := version.New(p.subject.Version())
	if err != nil {
		return false
	}

	return requested.Compare(actual) == 0
}

func (p *Platform) IsAndroidApp() bool {
	return p.IsAndroid() && regexp.MustCompile(`\bwv\b`).MatchString(p.userAgent)
}

func (p *Platform) IsAndroidWebview() bool {
	return p.IsAndroidApp()
}

func (p *Platform) IsBlackberry() bool {
	return p.id == "blackberry"
}

func (p *Platform) IsChromeOS() bool {
	return p.id == "chrome_os"
}

func (p *Platform) IsFirefoxOS() bool {
	return p.id == "firefox_os"
}

func (p *Platform) IsIOS() bool {
	return p.id == "ios"
}

func (p *Platform) IsIOSApp() bool {
	return p.IsIOS() && strings.Contains(p.userAgent, "Safari")
}

func (p *Platform) IsIOSWebview() bool {
	return p.IsIOSApp()
}

func (p *Platform) IsLinux() bool {
	return p.id == "linux"
}

func (p *Platform) IsMac() bool {
	return p.id == "mac"
}

func (p *Platform) IsMacV(expected string) bool {
	return p.IsMac() && util.DetectVersion(p.subject.Version(), expected)
}

func (p *Platform) IsUnknown() bool {
	return p.id == "unknown_platform"
}

func (p *Platform) IsWindows10() bool {
	return p.IsWindows() && strings.Contains(p.userAgent, "Windows NT 10")
}

func (p *Platform) IsWindows7() bool {
	return p.IsWindows() && strings.Contains(p.userAgent, "Windows NT 6.1")
}

func (p *Platform) IsWindows8() bool {
	return p.IsWindows() && regexp.MustCompile(`Windows NT 6\.[2-3]`).MatchString(p.userAgent)
}

func (p *Platform) IsWindows81() bool {
	return p.IsWindows() && strings.Contains(p.userAgent, "Windows NT 6.3")
}

func (p *Platform) IsWindows() bool {
	return p.id == "windows"
}

func (p *Platform) IsWindowsMobile() bool {
	return p.id == "windows_mobile"
}

func (p *Platform) IsWindowsPhone() bool {
	return p.id == "windows_phone"
}

func (p *Platform) IsWindowsRT() bool {
	return p.IsWindows8() && strings.Contains(p.userAgent, "ARM")
}

func (p *Platform) IsWindowsTouchscreenDesktop() bool {
	return p.IsWindows() && strings.Contains(p.userAgent, "Touch")
}

func (p *Platform) IsWindowsVista() bool {
	return p.IsWindows() && strings.Contains(p.userAgent, "Windows NT 6.0")
}

func (p *Platform) IsWindowsWOW64() bool {
	return p.IsWindows() && strings.Contains(p.userAgent, "WOW64")
}

func (p *Platform) IsWindowsX64() bool {
	return p.IsWindows() && regexp.MustCompile(`(Win64|x64|Windows NT 5\.2)`).MatchString(p.userAgent)
}

func (p *Platform) IsWindowsX64Inclusive() bool {
	return p.IsWindowsX64() || p.IsWindowsWOW64()
}

func (p *Platform) IsWindowsXP() bool {
	return p.IsWindows() && regexp.MustCompile(`Windows NT 5\.[12]`).MatchString(p.userAgent)
}

func (p *Platform) IsKaiOS() bool {
	return p.id == "kai_os"
}

func (p *Platform) ID() string {
	return p.id
}

func (p *Platform) Name() string {
	return p.name
}
