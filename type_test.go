// Copyright © 2026 Ruslan Sadekov.

package files

import (
	"encoding/json"
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

func TestType_UnmarshalJSON(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		wantErr bool
		json    any
		t       Type
	}{
		{
			name: "video string",
			json: "video",
			t:    Video,
		},
		{
			name: "Video string",
			json: "Video",
			t:    Video,
		},
		{
			name: "VIDEO string",
			json: "VIDEO",
			t:    Video,
		},
		{
			name: "AuDiO string",
			json: "AuDiO",
			t:    Audio,
		},
		{
			name: "audio string",
			json: "audio",
			t:    Audio,
		},
		{
			name: "image string",
			json: "image",
			t:    Image,
		},
		{
			name: "document string",
			json: "document",
			t:    Document,
		},
		{
			name: "archive string",
			json: "archive",
			t:    Archive,
		},
		{
			name: "font string",
			json: "font",
			t:    Font,
		},
		{
			name:    "unknown string",
			wantErr: true,
			json:    "unknown",
			t:       Unknown,
		},
		{
			name:    "invalid string",
			wantErr: true,
			json:    "invalid",
			t:       Unknown,
		},
		{
			name:    "too long string",
			wantErr: true,
			json:    strings.Repeat("1", maxTypeLen+1),
			t:       Unknown,
		},
		{
			name:    "too short string",
			wantErr: true,
			json:    strings.Repeat("1", minTypeLen-1),
			t:       Unknown,
		},
		{
			name: "video number",
			json: float64(Video),
			t:    Video,
		},
		{
			name: "audio number",
			json: float64(Audio),
			t:    Audio,
		},
		{
			name: "image number",
			json: float64(Image),
			t:    Image,
		},
		{
			name: "document number",
			json: float64(Document),
			t:    Document,
		},
		{
			name: "archive number",
			json: float64(Archive),
			t:    Archive,
		},
		{
			name: "font number",
			json: float64(Font),
			t:    Font,
		},
		{
			name: "unknown number",
			json: float64(Unknown),
			t:    Unknown,
		},
		{
			name:    "invalid number",
			wantErr: true,
			json:    float64(100),
			t:       Unknown,
		},
		{
			name: "zero number",
			json: float64(0),
			t:    Video,
		},
		{
			name:    "invalid type (bool)",
			wantErr: true,
			json:    true,
			t:       Unknown,
		},
		{
			name:    "negative number",
			wantErr: true,
			json:    float64(-100),
			t:       Unknown,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			data, err := json.Marshal(tt.json)
			if err != nil {
				t.Fatalf("failed to marshal input: %v", err)
			}

			got := Type(0)
			err = json.Unmarshal(data, &got)

			if (err != nil) != tt.wantErr {
				t.Errorf("Unmarshal() = %v, want: %t", err, tt.wantErr)
			}

			if got != tt.t {
				t.Errorf("want: %v, got: %v", tt.t, got)
			}
		})
	}
}

func TestType_UnmarshalText(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		wantErr bool
		text    string
		t       Type
	}{
		{
			name: "video string",
			text: "video",
			t:    Video,
		},
		{
			name: "Video string",
			text: "Video",
			t:    Video,
		},
		{
			name: "VIDEO string",
			text: "VIDEO",
			t:    Video,
		},
		{
			name: "AuDiO string",
			text: "AuDiO",
			t:    Audio,
		},
		{
			name: "audio string",
			text: "audio",
			t:    Audio,
		},
		{
			name: "image string",
			text: "image",
			t:    Image,
		},
		{
			name: "document string",
			text: "document",
			t:    Document,
		},
		{
			name: "archive string",
			text: "archive",
			t:    Archive,
		},
		{
			name: "font string",
			text: "font",
			t:    Font,
		},
		{
			name:    "unknown string",
			wantErr: true,
			text:    "unknown",
			t:       Unknown,
		},
		{
			name:    "invalid string",
			wantErr: true,
			text:    "invalid",
			t:       Unknown,
		},
		{
			name:    "too long string",
			wantErr: true,
			text:    strings.Repeat("1", maxTypeLen+1),
			t:       Unknown,
		},
		{
			name:    "too short string",
			wantErr: true,
			text:    strings.Repeat("1", minTypeLen-1),
			t:       Unknown,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := Type(0)
			err := got.UnmarshalText([]byte(tt.text))

			if (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalText() = %v, want: %t", err, tt.wantErr)
			}

			if got != tt.t {
				t.Errorf("want: %v, got: %v", tt.t, got)
			}
		})
	}
}

func TestType_MarshalText(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		wantErr bool
		t       Type
		want    string
	}{
		{name: "video", t: Video, want: "video"},
		{name: "audio", t: Audio, want: "audio"},
		{name: "image", t: Image, want: "image"},
		{name: "document", t: Document, want: "document"},
		{name: "archive", t: Archive, want: "archive"},
		{name: "font", t: Font, want: "font"},
		{name: "unknown", wantErr: true, t: Unknown},
		{name: "invalid", wantErr: true, t: Type(100)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//t.Parallel()

			got, err := tt.t.MarshalText()
			if (err != nil) != tt.wantErr {
				t.Errorf("MarshalText() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				return
			}

			if string(got) != tt.want {
				t.Errorf("MarshalText() = %v, want %v", got, tt.want)
			}
		})
	}
}
