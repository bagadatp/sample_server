package data

import (
	"net/url"
	"testing"
)

func TestGetBasicData(t *testing.T) {
	for _, tc := range []struct {
                name string
                q    url.Values
                e    string
	}{
		{
			"query empty",
			nil,
			"Response Default",
		},
		{
			"query a",
			map[string][]string{"a": {"abc"}},
			"Response A: abc",
		},
		{
			"query b",
			map[string][]string{"b": {"abc"}},
			"Response B: abc",
		},
		{
			"query c",
			map[string][]string{"c": {"abc"}},
			"Response C: abc",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			r := GetBasicData(tc.q)
			if r != tc.e {
				t.Fatalf("Unexpected result data generated: expected '%v' != received '%v'", tc.e, r)
			}
		})
	}
}