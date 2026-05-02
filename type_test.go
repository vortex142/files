// Copyright © 2026 Ruslan Sadekov.

package files

import (
	"strings"
	"testing"
)

func TestType_Validate(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		wantErr bool
		t       Type
	}{
		{name: "video", wantErr: false, t: Video},
		{name: "audio", wantErr: false, t: Audio},
		{name: "image", wantErr: false, t: Image},
		{name: "document", wantErr: false, t: Document},
		{name: "archive", wantErr: false, t: Archive},
		{name: "font", wantErr: false, t: Font},
		{name: "unknown", wantErr: false, t: Unknown},
		{name: "invalid", wantErr: true, t: Type(100)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			err := tt.t.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestType_String(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		t    Type
		want string
	}{
		{name: "video", t: Video, want: "video"},
		{name: "audio", t: Audio, want: "audio"},
		{name: "image", t: Image, want: "image"},
		{name: "document", t: Document, want: "document"},
		{name: "archive", t: Archive, want: "archive"},
		{name: "font", t: Font, want: "font"},
		{name: "unknown", t: Unknown, want: "unknown"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := tt.t.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestType_FromString(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		s    string
		want Type
	}{
		{name: "video", s: "video", want: Video},
		{name: "Video", s: "Video", want: Video},
		{name: "AUDIO", s: "AUDIO", want: Audio},
		{name: "audio", s: "audio", want: Audio},
		{name: "image", s: "image", want: Image},
		{name: "document", s: "document", want: Document},
		{name: "archive", s: "archive", want: Archive},
		{name: "font", s: "font", want: Font},
		{name: "unknown", s: "unknown", want: Unknown},
		{name: "invalid", s: "invalid", want: Unknown},
		{name: "too long string", s: strings.Repeat("1", maxTypeLen+1), want: Unknown},
		{name: "too short string", s: strings.Repeat("1", minTypeLen-1), want: Unknown},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := TypeFromString(tt.s); got != tt.want {
				t.Errorf("TypeFromString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestType_MinMaxConsts(t *testing.T) {
	var actualMaxLen int
	var actualMinLen int = 1000000

	for name := range typeNames {
		if len(name) > actualMaxLen {
			actualMaxLen = len(name)
		}

		if len(name) < actualMinLen {
			actualMinLen = len(name)
		}
	}

	if actualMaxLen != maxTypeLen {
		t.Errorf("actual max len = %d, current %d", actualMaxLen, maxTypeLen)
	}

	if actualMinLen != minTypeLen {
		t.Errorf("actual min len = %d, current %d", actualMinLen, minTypeLen)
	}
}
