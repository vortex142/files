// Copyright © 2026 Ruslan Sadekov.

package size

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"testing"
)

var parsePattern = regexp.MustCompile(`^([0-9.]+)\s*([a-zA-Z]+)$`)

func BenchmarkSize_Parse(b *testing.B) {
	needRegexp := false

	b.ResetTimer()
	b.ReportAllocs()

	benchs := []struct {
		name string
		str  string
	}{
		{
			name: "valid str",
			str:  "10 GB",
		},
		{
			name: "too long space",
			str:  "10                           GB",
		},
		{
			name: "low-case unit",
			str:  "10 gb",
		},
		{
			name: "empty str",
		},
		{
			name: "invalid unit",
			str:  "10 invalid",
		},
		{
			name: "only unit",
			str:  "GB",
		},
		{
			name: "only val",
			str:  "10",
		},
		{
			name: "reverse format",
			str:  "GB 10",
		},
		{
			name: "too long string",
			str:  strings.Repeat("55", maxParseLen),
		},
		{
			name: "num with space",
			str:  "10 000 Mb",
		},
		{
			name: "num with ','",
			str:  "1,5 Mb",
		},
	}

	if needRegexp {
		// regexp
		for _, bb := range benchs {
			b.Run(fmt.Sprintf("%s_regexp", bb.name), func(b *testing.B) {
				b.ResetTimer()
				for b.Loop() {
					regexpParseSize(bb.str)
				}
			})
		}
	}

	// optimize parser
	for _, bb := range benchs {
		b.Run(fmt.Sprintf("%s_optimize", bb.name), func(b *testing.B) {
			b.ResetTimer()
			for b.Loop() {
				s := Size(0)
				s.Parse(bb.str)
			}
		})
	}
}

func regexpParseSize(str string) (Size, error) {
	if str == "" {
		return 0, ErrEmptyParseStr
	}

	if len(str) > maxParseLen {
		return 0, ErrTooLongParseStr
	}

	// Normalization ensures compatibility with nameToUnit lookup regardless of user input style.
	str = strings.ToUpper(strings.TrimSpace(str))

	// The number of matches must be equal to 3, since the first element of the slice is a full string str.
	match := parsePattern.FindStringSubmatch(str)
	if len(match) != 3 {
		return 0, fmt.Errorf(`invalid size format: %s (expected e.g. "1.5 GB")`, str)
	}

	val, err := strconv.ParseFloat(match[1], 64)
	if err != nil {
		return 0, err
	}

	unitStr := match[2]
	unit, ok := nameToUnit[unitStr]
	if !ok {
		return 0, fmt.Errorf("unknown unit: %s", unitStr)
	}

	return From(val, unit), nil
}
