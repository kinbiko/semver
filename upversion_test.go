package semver_test

import (
	"fmt"
	"testing"

	"github.com/kinbiko/semver"
)

func TestUpversion(t *testing.T) {
	t.Run("invalid args", func(t *testing.T) {
		for _, tc := range []struct {
			current, step, exp string
		}{
			{exp: "no step value given"},
			{exp: "invalid step value 'breaking', valid values are ['major', 'minor', 'patch']", step: "breaking", current: "0.0.1"},
			{exp: "current version (fish) is not semver", step: "major", current: "fish"},
			{exp: "current version (a.1.1) is not semver", step: "major", current: "a.1.1"},
			{exp: "current version (1.b.1) is not semver", step: "major", current: "1.b.1"},
			{exp: "current version (1.1.c) is not semver", step: "major", current: "1.1.c"},
		} {
			t.Run(tc.exp, func(t *testing.T) {
				got, err := semver.Upversion(tc.step, tc.current)
				if got != "" {
					t.Errorf("expected no value return but got '%s'", got)
				}
				if err == nil {
					t.Fatal("expected non-nil error value")
				}
				if got := err.Error(); got != tc.exp {
					t.Errorf("expected error message '%s' but got '%s'", tc.exp, got)
				}
			})
		}
	})

	t.Run("positive test cases", func(t *testing.T) {
		for _, tc := range []struct {
			current, step, exp, errMsg string
		}{
			{current: "0.0.1", step: "patch", exp: "0.0.2"},
			{current: "0.1.1", step: "patch", exp: "0.1.2"},
			{current: "0.1.1", step: "minor", exp: "0.2.0"},
			{current: "1.1.1", step: "minor", exp: "1.2.0"},
			{current: "1.1.1", step: "major", exp: "2.0.0"},
			{current: "0.0.0", step: "major", exp: "1.0.0"},

			{current: "v0.0.1", step: "patch", exp: "v0.0.2"},
			{current: "v0.1.1", step: "patch", exp: "v0.1.2"},
			{current: "v0.1.1", step: "minor", exp: "v0.2.0"},
			{current: "v1.1.1", step: "minor", exp: "v1.2.0"},
			{current: "v1.1.1", step: "major", exp: "v2.0.0"},
			{current: "v0.0.0", step: "major", exp: "v1.0.0"},
		} {
			t.Run(fmt.Sprintf("%s upversion from '%s' should be '%s'", tc.step, tc.current, tc.exp), func(t *testing.T) {
				got, err := semver.Upversion(tc.step, tc.current)
				if err != nil {
					t.Fatal(err)
				}
				if got != tc.exp {
					t.Errorf("expected a %s upversion from '%s' to be '%s', but was '%s'", tc.step, tc.current, tc.exp, got)
				}
			})
		}
	})
}
