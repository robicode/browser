# Browser: Go Port of the Ruby Browser Gem

This is a straight Go port of the Ruby [Browser](https://rubygems.org/gems/browser) gem. Its purpose is to aid in getting various bits of information from the `UserAgent` header sent by browsers, search engines, etc.

I've tried to stay as true to the Ruby API as much as possible, while making concessions for Go best practices (like not using underscores in function names, for example). Most funcs use the IsXXX convention. Therefore, while in Ruby, you'd call `browser.bot?`, in Go you would call `browser.IsBot`.

Please note: I'm by no means an expert Gopher, so if you have any suggestions or contributions please feel free to open an issue.

## Unfinished Features

These features are unfinished and thus cannot be used:

* Accept-Language support
* tests

# Adapted README

What follows is the README from the Ruby library, adapted/modified for this Go library.

# Browser

*badges removed as not relevant for Go release*

Do some browser detection with Go.

## Installation

```bash
go get -u github.com/robicode/browser
```

## Usage

```go
import "github.com/robicode/browser"

browser := browser.New("path to browser yaml files", "Some User Agent", "en-us")

// General info
browser.IsBot()
browser.IsChrome()
browser.IsCoreMedia()
browser.IsDuckDuckGo()
browser.IsEdge()                // Newest MS browser
browser.IsElectron()            // Electron Framework
browser.IsFirefox()
browser.FullVersion()
browser.ID()
browser.IsIE()
browser.IsIE("6")            // detect specific IE version
browser.IsIE(">8", "<10")    // detect specific IE (IE9).
browser.IsKnown()            // has the browser been successfully detected?
browser.IsUnknown()          // the browser wasn't detected.
browser.Meta                 // an array with several attributes. Not implemented yet.
browser.Name()               // readable browser name
browser.IsNokia()
browser.IsOpera()
browser.IsOperaMini()
browser.IsPhantomJS()
browser.IsQuicktime()
browser.IsSafari()
browser.IsSafariWebappMode()
browser.IsSamsungBrowser()
browser.String()            // the meta info joined by space. Not implemented yet.
browser.IsUCBrowser()
browser.Version()           // major version number
browser.IsWebkit()
browser.WebkitFullVersion()
browser.IsYandex()
browser.IsWechat()
browser.IsQQ()
browser.IsWeibo()
browser.IsSputnik()
browser.IsSougouBrowser()

// Get bot info
browser.Bot.Name()
browser.Bot.IsSearchEngine()
browser.IsBot()
browser.Bot.Why() // shows which matcher detected this user agent as a bot.
// Not implemented yet - Browser::Bot.why?(ua)

// Get device info
browser.Device // (*Device)
browser.Device.ID()
browser.Device.Name()
browser.Device.IsUnknown()
browser.Device.IsblackberryPlaybook()
browser.Device.IsConsole()
browser.Device.IsIpad()
browser.Device.IsIphone()
browser.Device.IsIpodTouch()
browser.Device.IsKindle()
browser.Device.IsKindleFire()
browser.Device.IsMobile()
browser.Device.IsNintendo()
browser.Device.IsPlaystation()
browser.Device.IsPS3()
browser.Device.IsPS4()
browser.Device.IsPS5() // I added this but don't have a PS5 to test - please let me know
browser.Device.IsPSP()
browser.Device.IsSilk()
browser.Device.IsSurface()
browser.Device.IsTablet()
browser.Device.IsTV()
browser.device.IsVita()
browser.device.IsWii()
browser.device.IswiiU()
browser.device.IsSamsung()
browser.device.IsSwitch()
browser.device.IsXbox()
browser.device.IsXbox360()
browser.device.IsXboxOne()

// Get platform info
browser.Platform // (*Platform)
browser.Platform.ID()
browser.Platform.Name()
browser.Platform.Version()  // e.g. 9 (for iOS9)
browser.Platform.IsAdobeAir()
browser.Platform.IsAndroid()
browser.Platform.IsAndroid("4.2")   // detect Android Jelly Bean 4.2
browser.Platform.IsAndroidApp()     // detect webview in an Android app
browser.Platform.AndroidWebview() // alias for android_app?
browser.Platform.IsBlackberry()
browser.Platform.IsBlackberry("10") // detect specific BlackBerry version
browser.Platform.IsChromeOS()
browser.Platform.IsFirefoxOS()
browser.Platform.IsIOS()     // detect iOS
browser.Platform.IsIOS("9")  // detect specific iOS version
browser.Platform.IsIOSApp()     // detect webview in an iOS app
browser.Platform.IsIOSWebview() // alias for ios_app?
browser.Platform.IsLinux()
browser.Platform.IsMac()
browser.Platform.IsUnknown()
browser.Platform.IsWindows10()
browser.Platform.IsWindows7()
browser.Platform.IsWindows8()
browser.Platform.IsWindows81()
browser.Platform.IsWindows()
browser.Platform.IsWindowsMobile()
browser.Platform.IsWindowsPhone()
browser.Platform.IsWindowsRT()
browser.Platform.IsWindowsTouchscreenDesktop()
browser.Platform.IsWindowsVista()
browser.Platform.IsWindowsWOW64()
browser.Platform.IsWindowsX64()
browser.Platform.IsWindowsX64Inclusive()
browser.Platform.IsWindowsXP()
browser.Platform.IsKaiOS()
```

### Aliases

*This section is irrelevant for Go.*

### What's being detected?

- For a list of platform detections, check
  [lib/browser/platform.rb](https://github.com/fnando/browser/blob/master/lib/browser/platform.rb)
- For a list of device detections, check
  [lib/browser/device.rb](https://github.com/fnando/browser/blob/master/lib/browser/device.rb)
- For a list of bot detections, check
  [bots.yml](https://github.com/fnando/browser/blob/master/bots.yml)

### Detecting modern browsers

*TODO* convert to Go

### Rails integration

*Not relevant to Go*

### Accept Language

Parses the accept-language header from an HTTP request and produces an array of
language objects sorted by quality.

*TODO* This code, and the associated browser functionality, have not yet been converted to Go.

```ruby
browser = Browser.new("Some User Agent", accept_language: "en-us")

browser.accept_language.class
#=> Array

language = browser.accept_language.first

language.code
#=> "en"

language.region
#=> "US"

language.full
#=> "en-US"

language.quality
#=> 1.0

language.name
#=> "English/United States"
```

Result is always sorted in quality order from highest to lowest. As per the HTTP
spec:

- omitting the quality value implies 1.0.
- quality value equal to zero means that is not accepted by the client.

### Internet Explorer

Internet Explorer has a compatibility view mode that allows newer versions
(IE8+) to run as an older version. Browser will always return the navigator
version, ignoring the compatibility view version, when defined. If you need to
get the engine's version, you have to use `browse.MSIEVersion()` and
`browse.MSIEFullVersion()`.

So, let's say an user activates compatibility view in a IE11 browser. This is
what you'll get:

```go
browser.Version()
#=> 11

browser.FullVersion()
#=> 11.0

browser.MSIEVersion()
#=> 7

browser.MSIEFullVersion()
#=> 7.0

browser.IsCompatibilityView()
#=> true
```

This behavior changed in `v1.0.0`; previously there wasn't a way of getting the
real browser version.

### Safari

iOS webviews and web apps aren't detected as Safari anymore, so be aware of that
if that's your case. You can use a combination of platform and webkit detection
to do whatever you want.

```go
// iPad's Safari running as web app mode.
browser = browser.New(filepath.Join(app.HomeDirectory, "data", "browser"), "Mozilla/5.0 (iPad; U; CPU OS 3_2_1 like Mac OS X; en-us) AppleWebKit/531.21.10 (KHTML, like Gecko) Mobile/7B405")

browser.isSafari()
#=> false

browser.isWebkit()
#=> true

browser.Platform.IsIOS()
#=> true
```

### Bots

The bot detection is quite aggressive. Anything that matches at least one of the
following requirements will be considered a bot.

- Empty user agent string
- User agent that matches `/crawl|fetch|search|monitoring|spider|bot/`
- Any known bot listed under
  [bots.yml](https://github.com/fnando/browser/blob/master/bots.yml)

The following is irrelevant for Go:
To add custom matchers, you can add a callable object to
`Browser::Bot.matchers`. The following example matches everything that has a
`externalhit` substring on it. The bot name will always be `General Bot`.

```ruby
Browser::Bot.matchers << ->(ua, _browser) { ua.match?(/externalhit/i) }
```

To clear all matchers, including the ones that are bundled, use
`Browser::Bot.matchers.clear`. You can re-add built-in matchers by doing the
following:

```ruby
Browser::Bot.matchers += Browser::Bot.default_matchers
```

To restore v2's bot detection, remove the following matchers:

```ruby
Browser::Bot.matchers.delete(Browser::Bot::KeywordMatcher)
Browser::Bot.matchers.delete(Browser::Bot::EmptyUserAgentMatcher)
```

To add a matcher in Go, create a struct that follows the BotMatcher interface and add it to the `matchers` array
in `NewBot` in the `bot.go` file. For example, to create the above matcher in Go, you'd a file similar to the following:

```go
import ...

type ExternalHitMatcher struct {
  userAgent string
}

func newExternalHitMatcher(ua string) *ExternalHitMatcher {
  return &ExternalHitMatcher{
    userAgent: ua
  }
}

func (e *ExternalHitMatcher) Check() bool {
  return strings.Contains(strings.ToLower(e.userAgent), strings.ToLower("externalhit"))
}

func (e *ExternalHitMatcher) Name() string {
  return "External Hit Matcher"
}
```

### Middleware

*This section is irrelevant to the Go port.*

### Restrictions

- User agent has a size limit of 2048 bytes.
- Accept-Language has a size limit of 2048 bytes.

If size is not respected, then `browser.New()` will return an error.

## Development

### Versioning

This library follows http://semver.org.

### Writing code

Once you've made your great commits (include tests, please):

1. [Fork](http://help.github.com/forking/) browser
2. Create a topic branch - `git checkout -b my_branch`
3. Push to your branch - `git push origin my_branch`
4. Create a pull request
5. That's it!

Please respect the indentation rules and code style.

## Configuring environment

*Irrelevant to Go port.*

### Adding new features

Before using your time to code a new feature, open a ticket asking if it makes
sense and if it's on this project's scope.

Don't forget to add a new entry to `CHANGELOG.md`.

#### Adding a new bot

*TODO* work on the testing stuff

1. Add the user agent to `test/ua_bots.yml`.
2. Add the readable name to `bots.yml`. The key must be something that matches
   the user agent, in lowercased text.
3. Run tests.

Don't forget to add a new entry to `CHANGELOG.md`.

#### Adding a new search engine

*TODO* work on the testing stuff

1. Add the user agent to `test/ua_search_engines.yml`.
2. Add the same user agent to `test/ua_bots.yml`.
3. Add the readable name to `search_engines.yml`. The key must be something that
   matches the user agent, in lowercased text.
4. Run tests.

Don't forget to add a new entry to `CHANGELOG.md`.

#### Wrong browser/platform/device detection

If you know how to fix it, follow the "Writing code" above. Open an issue
otherwise

## Maintainer

- The maintainer of the original Rubygem is Nando Vieira - http://nandovieira.com
- The maintainer of this Go port is Joe Robison - https://robison.dev

## Contributors

- https://github.com/fnando/browser/contributors
- https://github.com/robicode/browser/contributors

## License

(The MIT License)

Please see the [LICENSE](https://github.com/robicode/browser/blob/main/LICENSE) file.
