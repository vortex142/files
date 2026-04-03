// Copyright © 2026 Ruslan Sadekov.

package mime

import (
	"strings"
	"testing"
)

func BenchmarkMime_Validation(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()

	benchs := []struct {
		name string
		m    Mime
	}{
		{
			name: "valid mime",
			m:    "audio/mpeg",
		},
		{
			name: "valid mime with optional",
			m:    "audio/mpeg;test=banana",
		},
		{
			name: "empty mime",
			m:    "",
		},
		{
			name: "invalid mime format",
			m:    "audio/mpeg/invalid",
		},
		{
			name: "empty type",
			m:    "/mpeg",
		},
		{
			name: "empty subtype",
			m:    "audio/",
		},
		{
			name: "empty type and subtype",
			m:    "/",
		},
		{
			name: "too long mime",
			m:    Mime(strings.Repeat("audio/mpeg", 1000000)),
		},
	}

	for _, bb := range benchs {
		b.Run(bb.name, func(b *testing.B) {
			b.ResetTimer()
			for b.Loop() {
				bb.m.Validate()
			}
		})
	}
}

func BenchmarkMime_Parts(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()

	benchs := []struct {
		name string
		m    Mime
	}{
		{
			name: "valid mime",
			m:    "audio/mpeg",
		},
		{
			name: "valid mime with optional",
			m:    "audio/mpeg;test=banana",
		},
		{
			name: "empty mime",
			m:    "",
		},
		{
			name: "invalid mime format",
			m:    "audio/mpeg/invalid",
		},
		{
			name: "empty type",
			m:    "/mpeg",
		},
		{
			name: "empty subtype",
			m:    "audio/",
		},
		{
			name: "empty type and subtype",
			m:    "/",
		},
		{
			name: "too long mime",
			m:    Mime(strings.Repeat("audio/mpeg", 1000000)),
		},
	}

	for _, bb := range benchs {
		b.Run(bb.name, func(b *testing.B) {
			b.ResetTimer()
			for b.Loop() {
				bb.m.Parts()
			}
		})
	}
}

func BenchmarkMime_PartsWithSplit(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()

	benchs := []struct {
		name string
		m    Mime
	}{
		{
			name: "valid mime",
			m:    "audio/mpeg",
		},
		{
			name: "valid mime with optional",
			m:    "audio/mpeg;test=banana",
		},
		{
			name: "empty mime",
			m:    "",
		},
		{
			name: "invalid mime format",
			m:    "audio/mpeg/invalid",
		},
		{
			name: "empty type",
			m:    "/mpeg",
		},
		{
			name: "empty subtype",
			m:    "audio/",
		},
		{
			name: "empty type and subtype",
			m:    "/",
		},
		{
			name: "too long mime",
			m:    Mime(strings.Repeat("audio/mpeg", 1000000)),
		},
	}

	for _, bb := range benchs {
		b.Run(bb.name, func(b *testing.B) {
			b.ResetTimer()
			for b.Loop() {
				mimePartsWithSplit(bb.m)
			}
		})
	}
}

func mimePartsWithSplit(m Mime) (string, Subtype, error) {
	if err := m.Validate(); err != nil {
		return "", "", err
	}

	parts := strings.Split(string(m), "/")

	// Strip any optional parameters (like "; charset=utf-8") to isolate the subtype.
	sub, _, _ := strings.Cut(parts[1], ";")

	return parts[0], Subtype(sub), nil
}
