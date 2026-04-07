// Copyright © 2026 Ruslan Sadekov.

package files

import (
	"strings"
	"testing"
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
			input:   Name(strings.Repeat("#", MaxNameLen)),
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
