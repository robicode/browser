package browser

// BrowserMatcher is the interface used to define a browser matcher.
// To add a new browser matcher, implement this interface and then add
// your matcher to the defaultMatchers array in browser.go.
type BrowserMatcher interface {
	FullVersion() string
	ID() string
	Matches() bool
	Name() string
}
