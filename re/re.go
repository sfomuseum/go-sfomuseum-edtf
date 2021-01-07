package re

import (
	"regexp"
)

const PATTERN_MY string = `^(?i)(c\.?\s+)?(\d+)/(\d{4}).*`

const PATTERN_MDY string = `^(?i)(c\.?\s+)?(\d+)/(\d+)/(\d{4}).*`

const PATTERN_EARLY string = `^(?i)(c\.?\s)?early\s+(\d{4})s.*`

const PATTERN_MID string = `^(?i)(c\.?\s+)?mid\s+(\d{4})s.*`

const PATTERN_LATE string = `^(?i)(c\.?\s+)?late\s+(\d{4})s.*`

const PATTERN_DECADE string = `^(?i)(c\.?\s+)?(\d{3})0s.*$`

const PATTERN_RANGE string = `^(?i)(c\.?\s+)?(\d{4})(?:\s+)?\-(?:\s+)?(\d{4}).*`

const PATTERN_YYYY string = `^(?i)(c\.?\s+)?(\s+\-\s+)?(\d{4}).*`

const PATTERN_MDYHUMAN string = `(?i)(\w+)\s+(\d{1,2})\,?\s+(\d{4}).*`

var MY *regexp.Regexp

var MDY *regexp.Regexp

var EARLY *regexp.Regexp

var MID *regexp.Regexp

var LATE *regexp.Regexp

var DECADE *regexp.Regexp

var RANGE *regexp.Regexp

var YYYY *regexp.Regexp

var MDYHUMAN *regexp.Regexp

func init() {

	MY = regexp.MustCompile(PATTERN_MY)

	MDY = regexp.MustCompile(PATTERN_MDY)

	EARLY = regexp.MustCompile(PATTERN_EARLY)

	MID = regexp.MustCompile(PATTERN_MID)

	LATE = regexp.MustCompile(PATTERN_LATE)

	DECADE = regexp.MustCompile(PATTERN_DECADE)

	RANGE = regexp.MustCompile(PATTERN_RANGE)

	YYYY = regexp.MustCompile(PATTERN_YYYY)

	MDYHUMAN = regexp.MustCompile(PATTERN_MDYHUMAN)
}
