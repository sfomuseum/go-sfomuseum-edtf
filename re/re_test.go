package re

import (
	"regexp"
	"testing"
)

func TestRegularExpressionPatterns(t *testing.T) {

	patterns := []string{
		PATTERN_MY,
		PATTERN_MY_LONGFORM,
		PATTERN_MDY,
		PATTERN_EARLY,
		PATTERN_MID,
		PATTERN_LATE,
		PATTERN_DECADE,
		PATTERN_RANGE,
		PATTERN_YYYY,
		PATTERN_MDYHUMAN,
	}

	for _, p := range patterns {

		t.Logf("Compile '%s'", p)

		_, err := regexp.Compile(p)

		if err != nil {

			t.Logf("Pattern '%s' failed to compile, %v", p, err)
			t.Fail()
			continue
		}

	}
}
