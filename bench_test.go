// Copyright © 2026 Ruslan Sadekov.

package files

import (
	"regexp"
	"strings"
	"testing"
)

var errResult error

var oldPattern = regexp.MustCompile(`^[ \p{L}\p{N}()._\-]+$`)

var (
	newlineRegex   = regexp.MustCompile(`[\n\r]`)
	forbiddenRegex = regexp.MustCompile(`[^a-zA-Zа-яА-Я0-9. ()_\-]`)
)

func benchsStruct() []struct {
	name string
	n    Name
} {
	return []struct {
		name string
		n    Name
	}{
		{
			name: "short name",
			n:    "sht",
		},
		{
			name: "medium name",
			n:    "banana-24.05.2022.mp4",
		},
		{
			name: "short win reserved name",
			n:    "con",
		},
		{
			name: "empty name",
		},
		{
			name: "long name",
			n:    Name(strings.Repeat("g", MaxNameLen)),
		},
		{
			name: "short name with invalid chars",
			n:    "sht:💋\n",
		},
		{
			name: "long name with invalid chars",
			n:    Name(strings.Repeat("#", MaxNameLen/2) + "1"),
		},
		{
			name: "long name with all chars is invalid",
			n:    Name(strings.Repeat("#", MaxNameLen/2)),
		},
	}
}

func BenchmarkName_Validate(b *testing.B) {
	b.ReportAllocs()

	needRegexp := false
	benchmarks := benchsStruct()

	if needRegexp {
		b.Run("validate_regexp", func(b *testing.B) {
			for _, bb := range benchmarks {
				b.ResetTimer()

				b.Run(bb.name, func(b *testing.B) {
					b.ResetTimer()
					for b.Loop() {
						benchmarkRegexp(string(bb.n))
					}
				})
			}
		})
	}

	b.Run("validate_native", func(b *testing.B) {
		for _, bb := range benchmarks {
			b.ResetTimer()

			b.Run(bb.name, func(b *testing.B) {
				b.ResetTimer()
				for b.Loop() {
					errResult = bb.n.Validate()
				}
			})
		}
	})
}

func BenchmarkName_Prepare(b *testing.B) {
	b.ReportAllocs()

	needRegexp := true
	benchs := benchsStruct()

	if needRegexp {
		b.Run("prepare_regexp", func(b *testing.B) {
			for _, bb := range benchs {
				initial := bb.n

				b.Run(bb.name, func(b *testing.B) {
					for b.Loop() {
						temp := initial

						prepareRegexp(string(temp), MaxNameLen)
					}
				})
			}
		})
	}

	b.Run("prepare_native", func(b *testing.B) {
		for _, bb := range benchs {
			initial := bb.n

			b.Run(bb.name, func(b *testing.B) {
				for b.Loop() {
					temp := initial

					temp.Sanitize()
				}
			})
		}
	})
}

func benchmarkRegexp(s string) bool {
	if s == "" {
		return false
	}

	if isReservedWinName(Name(s)) {
		return false
	}

	return oldPattern.MatchString(s)
}

func prepareRegexp(s string, maxLen int) string {
	if s == "" {
		return string("default name")
	}

	runes := []rune(s)
	if len(runes) > maxLen {
		runes = runes[:maxLen]
	}
	s = string(runes)

	s = newlineRegex.ReplaceAllString(s, "_")

	s = forbiddenRegex.ReplaceAllString(s, "")

	if len(s) == 0 {
		return string("default name")
	}

	if isReservedWinName(Name(s)) {
		return "_" + s
	}

	return s
}
