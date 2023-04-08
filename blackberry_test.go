package browser

import "testing"

type Test struct {
	UserAgent           string
	ExpectedVersion     string
	ExpectedFullVersion string
	ExpectedMatches     bool
	Name                string
}

var tests = []Test{
	{
		UserAgent:           "Mozilla/5.0 (BB10; Touch) AppleWebKit/537.10+ (KHTML, like Gecko) Version/10.0.9.1675 Mobile Safari/537.10+",
		ExpectedVersion:     "10",
		ExpectedFullVersion: "10.0.9.1675",
		ExpectedMatches:     true,
		Name:                "Blackberry 10",
	},
	{
		UserAgent:           "BlackBerry8100/4.2.1 Profile/MIDP-2.0 Configuration/CLDC-1.1 VendorID/103",
		ExpectedVersion:     "4",
		ExpectedFullVersion: "4.2.1",
		ExpectedMatches:     true,
		Name:                "Blackberry 4",
	},
}

func Test_Blackberry(t *testing.T) {
	for _, test := range tests {
		bb := newBlackberry(test.UserAgent)

		if bb.Matches() != test.ExpectedMatches {
			t.Error("expected test", test.Name, "to be", test.ExpectedMatches, "but was", bb.Matches())
		}

		if bb.FullVersion() != test.ExpectedFullVersion {
			t.Error("expected test", test.Name, "to return FullVersion() of", test.ExpectedFullVersion, "but was", bb.FullVersion())
		}
	}
}
