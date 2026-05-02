// Copyright © 2026 Ruslan Sadekov.

package files

import (
	"strings"
	"testing"
)

func typeFromString_WithOnlyToLower(s string) Type {
	s = strings.ToLower(s)
	t, found := typeNames[s]
	if !found {
		return Unknown
	}
	return t
}

func typeFromString_WithLengthCheck(s string) Type {
	if len(s) > maxTypeLen || len(s) < minTypeLen {
		return Unknown
	}

	return typeFromString_WithOnlyToLower(s)
}

func typeFromString_WithLazyToLower(s string) Type {
	if len(s) > maxTypeLen || len(s) < minTypeLen {
		return Unknown
	}

	t, found := typeNames[s]
	if !found {
		s = strings.ToLower(s)
		t, found = typeNames[s]
	}

	if !found {
		return Unknown
	}
	return t
}

func BenchmarkType_FromString(b *testing.B) {
	b.ReportAllocs()

	benchs := []struct {
		name string
		s    string
	}{
		{name: "video", s: "video"},
		{name: "Video", s: "Video"},
		{name: "audio", s: "audio"},
		{name: "AUDIO", s: "AUDIO"},
		{name: "image", s: "image"},
		{name: "document", s: "document"},
		{name: "archive", s: "archive"},
		{name: "font", s: "font"},
		{name: "unknown", s: "unknown"},
		{name: "empty", s: ""},
		{name: "invalid (lowercase)", s: "invalid"},
		{name: "invalid (uppercase)", s: "INVALID"},
		{name: "long", s: strings.Repeat("g", maxTypeLen)},
		{name: "short", s: strings.Repeat("g", minTypeLen-1)},
	}

	needOnlyToLowerBenchmark := false
	needLengthCheckBenchmark := false
	needLazyToLowerBenchmark := false

	if needOnlyToLowerBenchmark {
		b.Run("only_to_lower", func(b *testing.B) {
			for _, bb := range benchs {
				b.ResetTimer()

				b.Run(bb.name, func(b *testing.B) {
					b.ResetTimer()
					for b.Loop() {
						typeFromString_WithOnlyToLower(bb.s)
					}
				})
			}
		})
	}

	if needLengthCheckBenchmark {
		b.Run("length_check", func(b *testing.B) {
			for _, bb := range benchs {
				b.ResetTimer()

				b.Run(bb.name, func(b *testing.B) {
					b.ResetTimer()
					for b.Loop() {
						typeFromString_WithLengthCheck(bb.s)
					}
				})
			}
		})
	}

	if needLazyToLowerBenchmark {
		b.Run("lazy_to_lower", func(b *testing.B) {
			for _, bb := range benchs {
				b.ResetTimer()

				b.Run(bb.name, func(b *testing.B) {
					b.ResetTimer()
					for b.Loop() {
						typeFromString_WithLazyToLower(bb.s)
					}
				})
			}
		})
	}

	b.Run("current", func(b *testing.B) {
		for _, bb := range benchs {
			b.ResetTimer()

			b.Run(bb.name, func(b *testing.B) {
				b.ResetTimer()
				for b.Loop() {
					TypeFromString(bb.s)
				}
			})
		}
	})
}
