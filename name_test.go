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

func TestAllowed(t *testing.T) {
	tests := []struct {
		name        string
		wantAllowed bool
		rs          []rune
	}{
		{
			name:        "allowed en rune",
			wantAllowed: true,
			rs:          []rune{'R', 'U', 'S', 'L', 'A', 'N'},
		},
		{
			name:        "allowed ru rune",
			wantAllowed: true,
			rs:          []rune{'Р', 'У', 'С', 'Л', 'А', 'Н'},
		},
		{
			name:        "allowed num",
			wantAllowed: true,
			rs:          []rune{'0', '1', '2', '3', '4'},
		},
		{
			name:        "allowed spec chars",
			wantAllowed: true,
			rs:          []rune{'(', ')', '.', '_', '-', ' '},
		},
		{
			name:        "not allowed control chars",
			wantAllowed: false,
			rs:          []rune{'\n', '\r', '\t'},
		},
		{
			name:        "not allowed char",
			wantAllowed: false,
			rs:          []rune{'|', ':', '+'},
		},
		{
			name:        "not allowed emoji",
			wantAllowed: false,
			rs:          []rune{'🚀', '😊'},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, r := range tt.rs {
				got := isAllowed(r)
				if got != tt.wantAllowed {
					t.Errorf("isAllowed(%c) [%U] = %v; want %v", r, r, got, tt.wantAllowed)
				}
			}
		})
	}
}

func TestName_Validate(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
		n       Name
	}{
		{
			name:    "valid name",
			wantErr: false,
			n:       "BANANA.MP4",
		},
		{
			name:    "forbidden windows name",
			wantErr: true,
			n:       "con",
		},
		{
			name:    "empty name",
			wantErr: true,
		},
		{
			name:    "valid name with chars",
			wantErr: false,
			n:       "BANANA_SUPER_FRUIT-KAIF(APPLE).MP4",
		},
		{
			name:    "too long name",
			wantErr: true,
			n:       Name(strings.Repeat("a", MaxNameLen+1)),
		},
		{
			name:    "contains invalid chars",
			wantErr: true,
			n:       "BANANA:SUPER:FRUIT|KAIF(APPLE).MP4",
		},
		{
			name:    "try path traversal inject",
			wantErr: true,
			n:       "../../etc/passwords",
		},
		{
			name:    "contains emoji",
			wantErr: true,
			n:       "РАКЕТА🚀.mp4",
		},
		{
			name:    `contans '\n'`,
			wantErr: true,
			n:       "BANANA_SUPER_FRUIT-KAIF(APPLE).MP4\nРАКЕТА!",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.n.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestName_CutExtension(t *testing.T) {
	tests := []struct {
		name          string
		input         Name
		want          Name
		wantExtension Extension
	}{
		{
			name:  "cut with zero dots",
			input: "BANANA",
			want:  "BANANA",
		},
		{
			name:          "cut with one dot",
			input:         "BANANA.mp3",
			want:          "BANANA",
			wantExtension: "mp3",
		},
		{
			name:          "cut with two dots",
			input:         "BANANA.1.mp3",
			want:          "BANANA.1",
			wantExtension: "mp3",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ext := tt.input.CutExtension()
			if tt.input != tt.want {
				t.Errorf("CutExtension() = %s, want %s", tt.input, tt.want)
				return
			}

			if tt.wantExtension != ext {
				t.Errorf("want: %s, got: %s", tt.wantExtension, ext)
			}
		})
	}
}

func TestName_Prepare(t *testing.T) {
	tests := []struct {
		name        string
		wantErr     bool
		input       Name
		want        Name
		defaultName Name
	}{
		{
			name:    "replace newlines and clean symbols",
			wantErr: false,
			input:   "Cool\nVideo! @🚀",
			want:    "Cool_Video ",
		},
		{
			name:    "empty input",
			wantErr: true,
		},
		{
			name:    "forbidden windows name",
			wantErr: false,
			input:   "con",
			want:    "_con",
		},
		{
			name:    "trim long cyrillic caption",
			wantErr: false,
			input:   Name(strings.Repeat("Я", MaxNameLen+10)),
			want:    Name(strings.Repeat("Я", MaxNameLen)),
		},
		{
			name:    "all forbidden chars",
			wantErr: true,
			input:   "!!!@@@###",
		},
		{
			name:    "mixed content with brackets",
			wantErr: false,
			input:   "Movie (2024).mp4\r\n[4K]",
			want:    "Movie (2024).mp4__4K",
		},
		{
			name:    "replace newlines and clean symbols",
			wantErr: false,
			input:   "Cool\nVideo! @🚀",
			want:    "Cool_Video ",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := tt.input
			err := c.Prepare()
			if (err != nil) != tt.wantErr {
				t.Errorf("Prepare() = %q, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				return
			}

			if c != tt.want {
				t.Errorf("Prepare() = %q, want %q", c, tt.want)
			}
		})
	}
}

func TestReservedWinNames(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"CON", true},
		{"con", true},
		{"PRN", true},
		{"aux", true},
		{"Nul", true},
		{"COM1", true},
		{"LPT9", true},

		{"", false},
		{"C", false},
		{"CO", false},
		{"CONN", false},
		{"CON1", false},
		{"BANANA", false},
		{"COM10", false},
		{"LPT0", false},
		{"_CON", false},
		{".AUX", false},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if got := isReservedWinName(Name(tt.input)); got != tt.want {
				t.Errorf("isReservedWinName(%q) = %v; want %v", tt.input, got, tt.want)
			}
		})
	}
}

func BenchmarkName_Validate(b *testing.B) {
	benchmarks := []struct {
		name string
		n    Name
	}{
		{
			name: "short name",
			n:    "sht",
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
			n:    Name(strings.Repeat("1💋", MaxNameLen/2-1)),
		},
	}

	for _, bb := range benchmarks {
		b.Run(bb.name, func(b *testing.B) {
			b.ReportAllocs()
			b.ResetTimer()

			for b.Loop() {
				errResult = bb.n.Validate()
			}
		})
	}
}

func BenchmarkValidate_Regexp(b *testing.B) {
	benchmarks := []struct {
		name string
		n    Name
	}{
		{
			name: "short name",
			n:    "sht",
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
			n:    Name(strings.Repeat("1💋", MaxNameLen/2-1)),
		},
	}

	for _, bb := range benchmarks {
		b.Run(bb.name, func(b *testing.B) {
			b.ReportAllocs()
			b.ResetTimer()

			for b.Loop() {
				benchmarkRegexp(string(bb.n))
			}
		})
	}
}

func BenchmarkValidate_Builder(b *testing.B) {
	input := Name("BANANA_SUPER_FRUIT-KAIF(APPLE).MP4")
	for b.Loop() {
		input.Validate()
	}
}

func BenchmarkPrepare_Regexp(b *testing.B) {
	input := "Cool\nVideo! @🚀 Очень длинное название видео для проверки лимитов. Настолько длинное что оно даже не должно поместиться полность в экранное пространство моего IDE (VS CODE): И Я ДОБИЛСЯ ЭТОГО"
	for b.Loop() {
		prepareRegexp(input, MaxNameLen)
	}
}

func BenchmarkPrepare_Builder(b *testing.B) {
	input := Name("Cool\nVideo! @🚀 Очень длинное название видео для проверки лимитов. Настолько длинное что оно даже не должно поместиться полность в экранное пространство моего IDE (VS CODE): И Я ДОБИЛСЯ ЭТОГО")
	for b.Loop() {
		input.Prepare()
	}
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
