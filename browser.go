package browser

import (
	"errors"
	"regexp"
	"strings"

	"github.com/robicode/browser/util"
)

type Browser struct {
	fullVersion string
	id          string
	matchers    []BrowserMatcher
	name        string
	ua          string

	Bot      *Bot
	Device   *Device
	Platform *Platform
	Meta     []string
	subject  BrowserMatcher
}

const (
	AcceptLanguageMaxLength int = 2048
	UserAgentMaxLength      int = 2048
)

// New creates a newly initialized Browser struct given the path
// to the YAML files, the user agent and an optional Accept-Languages header.
func New(dataPath, userAgent, acceptLanguage string) (*Browser, error) {
	if len(userAgent) > UserAgentMaxLength {
		return nil, errors.New("user agent too long")
	}

	if len(acceptLanguage) > AcceptLanguageMaxLength {
		return nil, errors.New("Accept-Language header too long")
	}

	bot, err := NewBot(dataPath, userAgent)
	if err != nil {
		return nil, err
	}

	device := NewDevice(dataPath, userAgent)
	platform := NewPlatform(userAgent)

	browser := &Browser{
		Bot:      bot,
		Device:   device,
		Platform: platform,
		ua:       userAgent,
	}

	nokia := newNokia(userAgent)
	ucbrowser := newUCBrowser(userAgent)
	phantomjs := newPhantomJS(userAgent)
	blackberry := newBlackberry(userAgent)
	opera := newOpera(userAgent)
	edge := newEdge(userAgent)
	ie := newInternetExplorer(userAgent)
	ff := newFirefox(userAgent)
	otter := newOtter(userAgent)
	fb := newFacebook(userAgent)
	instagram := newInstagram(userAgent)
	snapchat := newSnapchat(userAgent)
	weibo := newWeibo(userAgent)
	mm := newMicroMessenger(userAgent)
	qq := newQQ(userAgent)
	alipay := newAlipayMatcher(userAgent)
	electron := newElectron(userAgent)
	yandex := newYandex(userAgent)
	sputnik := newSputnik(userAgent)
	ddg := newDuckDuckGo(userAgent)
	samsung := newSamsungBrowser(userAgent)
	huawei := newHuaweiBrowser(userAgent)
	miui := newMiuiBrowser(userAgent)
	maxthon := newMaxthon(userAgent)
	sougu := newSougouBrowser(userAgent)
	gsa := newGoogleSearchApp(userAgent, browser)
	chrome := newChromeMatcher(userAgent, browser)
	safari := newSafari(userAgent, browser)
	unknown := newUnknown(userAgent)

	defaultMatchers := []BrowserMatcher{
		nokia,
		ucbrowser,
		phantomjs,
		blackberry,
		opera,
		edge,
		ie,
		ff,
		otter,
		fb,
		instagram,
		snapchat,
		weibo,
		mm,
		qq,
		alipay,
		electron,
		yandex,
		sputnik,
		ddg,
		samsung,
		huawei,
		miui,
		maxthon,
		sougu,
		gsa,
		chrome,
		safari,
		unknown,
	}

	browser.matchers = defaultMatchers

	for _, matcher := range browser.matchers {
		if matcher.Matches() {
			browser.subject = matcher
			browser.fullVersion = matcher.FullVersion()
			browser.id = matcher.ID()
			browser.name = matcher.Name()
			break
		}
	}

	return browser, nil
}

// IsSnapchat detects if the browser is Snapchat
func (b *Browser) IsSnapchat(expectedVersion ...string) bool {
	if len(expectedVersion) > 0 {
		return newSnapchat(b.ua).Matches() && util.DetectVersion(b.subject.FullVersion(), expectedVersion)
	}

	return newSnapchat(b.ua).Matches()
}

func (b *Browser) IsBlackberry(expectedVersion ...string) bool {
	if len(expectedVersion) > 0 {
		return newBlackberry(b.ua).Matches() && util.DetectVersion(b.subject.FullVersion(), expectedVersion)
	}

	return newBlackberry(b.ua).Matches()
}

// IsOtter detects if the browser is Otter
func (b *Browser) IsOtter(expectedVersion ...string) bool {
	if len(expectedVersion) > 0 {
		return newOtter(b.ua).Matches() && util.DetectVersion(b.subject.FullVersion(), expectedVersion)
	}

	return newOtter(b.ua).Matches()
}

// IsBot detects if the browser is a bot
func (b *Browser) IsBot() bool {
	return b.Bot.IsBot()
}

// IsHuaweiBrowser detects if the browser is Huawei Browser
func (b *Browser) IsHuaweiBrowser(expectedVersion ...string) bool {
	if len(expectedVersion) > 0 {
		return newHuaweiBrowser(b.ua).Matches() && util.DetectVersion(b.subject.FullVersion(), expectedVersion)
	}

	return newHuaweiBrowser(b.ua).Matches()
}

// IsAlipay detects if the browser is Alipay
func (b *Browser) IsAlipay(expectedVersion ...string) bool {
	if len(expectedVersion) > 0 {
		return newAlipayMatcher(b.ua).Matches() && util.DetectVersion(b.subject.FullVersion(), expectedVersion)
	}

	return newAlipayMatcher(b.ua).Matches()
}

// IsMiuiBrowser detects if the browser is Miui Browser
func (b *Browser) IsMiuiBrowser(expectedVersion ...string) bool {
	if len(expectedVersion) > 0 {
		return newMiuiBrowser(b.ua).Matches() && util.DetectVersion(b.subject.FullVersion(), expectedVersion)
	}

	return newMiuiBrowser(b.ua).Matches()
}

// IsMicroMessenger detects if the browser is Micro Messenger
func (b *Browser) IsMicroMessenger(expectedVersion ...string) bool {
	if len(expectedVersion) > 0 {
		return newMicroMessenger(b.ua).Matches() && util.DetectVersion(b.subject.FullVersion(), expectedVersion)
	}

	return newMicroMessenger(b.ua).Matches()
}

// IsMaxthon detects if the browser is Maxthon
func (b *Browser) IsMaxthon(expectedVersion ...string) bool {
	if len(expectedVersion) > 0 {
		return newMaxthon(b.ua).Matches() && util.DetectVersion(b.subject.FullVersion(), expectedVersion)
	}

	return newMaxthon(b.ua).Matches()
}

// IsSougouBrowser detects if browser is Sougou
func (b *Browser) IsSougouBrowser(expectedVersion ...string) bool {
	if len(expectedVersion) > 0 {
		return newSougouBrowser(b.ua).Matches() && util.DetectVersion(b.subject.FullVersion(), expectedVersion)
	}

	return newSougouBrowser(b.ua).Matches()
}

// IsGoogleSearchApp detects if browser is GOogle Search App
func (b *Browser) IsGoogleSearchApp(expectedVersion ...string) bool {
	if len(expectedVersion) > 0 {
		return newGoogleSearchApp(b.ua, b).Matches() && util.DetectVersion(b.subject.FullVersion(), expectedVersion)
	}

	return newGoogleSearchApp(b.ua, b).Matches()
}

// IsChrome detects if the browser is Google Chrome
func (b *Browser) IsChrome(expectedVersion ...string) bool {
	if len(expectedVersion) > 0 {
		return newChromeMatcher(b.ua, b).Matches() && util.DetectVersion(b.subject.FullVersion(), expectedVersion)
	}

	return newChromeMatcher(b.ua, b).Matches()
}

// IsCoreMedia Detects if browser is Apple CoreMedia
func (b *Browser) IsCoreMedia(expectedVersion ...string) bool {
	if len(expectedVersion) > 0 {
		return strings.Contains(b.ua, "CoreMedia") &&
			util.DetectVersion(b.subject.FullVersion(), expectedVersion)
	}

	return strings.Contains(b.ua, "CoreMedia")
}

// IsDuckDuckGo detects if the browser is Duck Duck Go
func (b *Browser) IsDuckDuckGo(expectedVersion ...string) bool {
	if len(expectedVersion) > 0 {
		return strings.Contains(b.ua, "DuckDuckGo") && util.DetectVersion(b.subject.FullVersion(), expectedVersion)
	}

	return strings.Contains(b.ua, "DuckDuckGo")
}

// IsEdge detects if the browser is Microsoft Edge
func (b *Browser) IsEdge(expectedVersion ...string) bool {
	if len(expectedVersion) > 0 {
		return newEdge(b.ua).Matches() && util.DetectVersion(b.Version(), expectedVersion)
	}

	return newEdge(b.ua).Matches()
}

// IsInstagram detects if the browser is Instagram
func (b *Browser) IsInstagram(expectedVersion ...string) bool {
	if len(expectedVersion) > 0 {
		return newInstagram(b.ua).Matches() && util.DetectVersion(b.Version(), expectedVersion)
	}

	return newInstagram(b.ua).Matches()
}

// IsElectron detects if the browser is Electron
func (b *Browser) IsElectron(expectedVersion ...string) bool {
	if len(expectedVersion) > 0 {
		return newElectron(b.ua).Matches() && util.DetectVersion(b.subject.FullVersion(), expectedVersion)
	}

	return newElectron(b.ua).Matches()
}

// IsFacebook detects whether the browser is Facebook
func (b *Browser) IsFacebook(expectedVersion ...string) bool {
	if len(expectedVersion) > 0 {
		return newFacebook(b.ua).Matches() && util.DetectVersion(b.Version(), expectedVersion)
	}

	return newFacebook(b.ua).Matches()
}

// IsFirefox detects if the browser is Firefox
func (b *Browser) IsFirefox(expectedVersion ...string) bool {
	if len(expectedVersion) > 0 {
		return newFirefox(b.ua).Matches() && util.DetectVersion(b.subject.FullVersion(), expectedVersion)
	}

	return newFirefox(b.ua).Matches()
}

func (b *Browser) FullVersion() string {
	return b.subject.FullVersion()
}

// IsIE detects if the browser is Internet Explorer
func (b *Browser) IsIE(expectedVersion ...string) bool {
	if len(expectedVersion) > 0 {
		return newInternetExplorer(b.ua).Matches() && util.DetectVersion(b.Version(), expectedVersion)
	}

	return newInternetExplorer(b.ua).Matches()
}

func (b *Browser) IsKnown() bool {
	return !b.IsUnknown()
}

func (b *Browser) IsUnknown() bool {
	return b.id == "unknown_browser"
}

// Name returns the browser name
func (b *Browser) Name() string {
	return b.subject.Name()
}

// IsNokia detects if the browser is Nokia
func (b *Browser) IsNokia(expectedVersion ...string) bool {
	if len(expectedVersion) > 0 {
		return newNokia(b.ua).Matches() && util.DetectVersion(b.subject.FullVersion(), expectedVersion)
	}

	return newNokia(b.ua).Matches()
}

// IsOpera detects if the browser is Opera
func (b *Browser) IsOpera(expectedVersion ...string) bool {
	if len(expectedVersion) > 0 {
		return newOpera(b.ua).Matches() && util.DetectVersion(b.subject.FullVersion(), expectedVersion)
	}

	return newOpera(b.ua).Matches()
}

// IsOperaMini detects if the browser is Opera Mini
func (b *Browser) IsOperaMini(expectedVersion ...string) bool {
	if len(expectedVersion) > 0 {
		return strings.Contains(b.ua, "Opera Mini") && util.DetectVersion(b.subject.FullVersion(), expectedVersion)
	}

	return strings.Contains(b.ua, "Opera Mini")
}

// IsPhantomJS detects if the browser is PhantomJS
func (b *Browser) IsPhantomJS(expectedVersion ...string) bool {
	if len(expectedVersion) > 0 {
		return newPhantomJS(b.ua).Matches() && util.DetectVersion(b.subject.FullVersion(), expectedVersion)
	}

	return newPhantomJS(b.ua).Matches()
}

// IsQuickTime detects whether browser is QuickTime
func (b *Browser) IsQuicktime(expectedVersion ...string) bool {
	if len(expectedVersion) > 0 {
		return strings.Contains(strings.ToLower(b.ua), strings.ToLower("IsQuickTime")) &&
			util.DetectVersion(b.subject.FullVersion(), expectedVersion)
	}

	return strings.Contains(strings.ToLower(b.ua), strings.ToLower("IsQuickTime"))
}

// IsSafari detects if the browser is Safari
func (b *Browser) IsSafari(expectedVersion ...string) bool {
	if len(expectedVersion) > 0 {
		return newSafari(b.ua, b).Matches() && util.DetectVersion(b.subject.FullVersion(), expectedVersion)
	}

	return newSafari(b.ua, b).Matches()
}

// IsSafariWebappMode detects if browser is Safari running in an iOS application
func (b *Browser) IsSafariWebappMode() bool {
	return (b.Device.IsIpad() || b.Device.IsIphone()) && strings.Contains(b.ua, "AppleWebKit")
}

// IsSamsung detects if the browser is Samsung
func (b *Browser) IsSamsungBrowser(expectedVersion ...string) bool {
	if len(expectedVersion) > 0 {
		return strings.Contains(b.ua, "SamsungBrowser") && util.DetectVersion(b.subject.FullVersion(), expectedVersion)
	}

	return strings.Contains(b.ua, "SamsungBrowser")
}

func (b *Browser) String() string {
	return ""
}

// IsUCBrowser detects if the browser is UCBrowser
func (b *Browser) IsUcBrowser(expectedVersion ...string) bool {
	if len(expectedVersion) > 0 {
		return newUCBrowser(b.ua).Matches() && util.DetectVersion(b.subject.FullVersion(), expectedVersion)
	}

	return newUCBrowser(b.ua).Matches()
}

// Version returns the browser version
func (b *Browser) Version() string {
	return strings.Split(b.subject.FullVersion(), ".")[0]
}

// IsWebkit detects if the browser is Webkit based
func (b *Browser) IsWebkit(expectedVersion string) bool {
	if len(expectedVersion) > 0 {
		return strings.Contains(b.ua, "AppleWebKit") &&
			(!b.IsEdge() || newEdge(b.ua).IsChromeBased()) &&
			util.DetectVersion(b.Version(), expectedVersion)
	}

	return strings.Contains(b.ua, "AppleWebKit") &&
		(!b.IsEdge() || newEdge(b.ua).IsChromeBased())
}

// IsYandex detects if the browser is Yandex
func (b *Browser) IsYandex(expectedVersion ...string) bool {
	if len(expectedVersion) > 0 {
		return newYandex(b.ua).Matches() && util.DetectVersion(b.subject.FullVersion(), expectedVersion)
	}

	return newYandex(b.ua).Matches()
}

// IsSputnik detects if the browser is Sputnik
func (b *Browser) IsSputnik(expectedVersion ...string) bool {
	if len(expectedVersion) > 0 {
		return newSputnik(b.ua).Matches() && util.DetectVersion(b.subject.FullVersion(), expectedVersion)
	}

	return newSputnik(b.ua).Matches()
}

// IsWechat is an alias for IsMicroMessenger
func (b *Browser) IsWechat(expectedVersion ...string) bool {
	return b.IsMicroMessenger(expectedVersion...)
}

// IsQQ detects if browser is QQ
func (b *Browser) IsQQ(expectedVersion ...string) bool {
	if len(expectedVersion) > 0 {
		return newQQ(b.ua).Matches() && util.DetectVersion(b.subject.FullVersion(), expectedVersion)
	}

	return newQQ(b.ua).Matches()
}

// IsWeibo detects if the browser is Weibo
func (b *Browser) IsWeibo(expectedVersion ...string) bool {
	if len(expectedVersion) > 0 {
		return newWeibo(b.ua).Matches() && util.DetectVersion(b.subject.FullVersion(), expectedVersion)
	}

	return newWeibo(b.ua).Matches()
}

// WebkitFullVersion returns the version of Webkit in use or "0.0"
func (b *Browser) WebkitFullVersion() string {
	matches := regexp.MustCompile(`AppleWebKit/([\d.]+)`).FindStringSubmatch(b.ua)
	if len(matches) > 0 {
		return matches[1]
	}

	return "0.0"
}

// IsProxy detects if the browser is a proxy browser
func (b *Browser) IsProxy() bool {
	return b.IsNokia() || b.IsUcBrowser() || b.IsOperaMini()
}

// IsCompatibilityView detects if IE is running in Compatibility View
func (b *Browser) IsCompatibilityView() bool {
	ie := newInternetExplorer(b.ua)
	if !ie.Matches() {
		return false
	}

	return ie.IsCompatibilityView()
}

// MSIEVersion detects the IE version of browser
func (b *Browser) MSIEVersion() string {
	ie := newInternetExplorer(b.ua)
	if !ie.Matches() {
		return "0"
	}

	return ie.MSIEVersion()
}

// MSIEFullVersion returns the full version of IE in use
func (b *Browser) MSIEFullVersion() string {
	ie := newInternetExplorer(b.ua)
	if !ie.Matches() {
		return "0.0"
	}

	return ie.MSIEFullVersion()
}
