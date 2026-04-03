// Copyright © 2026 Ruslan Sadekov.

package mime

import "testing"

func TestSubtype_Name(t *testing.T) {
	tests := []struct {
		name    string
		subtype Subtype
		want    string
		wantErr bool
	}{
		{"vnd.wave", "vnd.wave", wavName, false},
		{"x-wav", "x-wav", wavName, false},
		{"mpeg", "mpeg", mpegName, false},
		{"mp3", "mp3", mpegName, false},
		{"mp4", "mp4", videoType, false},
		{"ogg", "ogg", oggName, false},
		{"unknown", "unknown", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.subtype.Name()
			if (err != nil) != tt.wantErr {
				t.Errorf("Name(%s) = %s, err: %v, wantErr %t", tt.subtype, got, err, tt.wantErr)
			}
		})
	}
}

func TestSubtype_IsWav(t *testing.T) {
	tests := []struct {
		name string
		sub  Subtype
		want bool
	}{
		{
			name: "valid vnd.wave",
			sub:  "vnd.wave",
			want: true,
		},
		{
			name: "valid vnd.wav",
			sub:  "vnd.wav",
			want: true,
		},
		{
			name: "valid wave",
			sub:  "wave",
			want: true,
		},
		{
			name: "valid x-wav",
			sub:  "x-wav",
			want: true,
		},
		{
			name: "valid x-wave",
			sub:  "x-wave",
			want: true,
		},
		{
			name: "invalid wav extension",
			sub:  "json",
			want: false,
		},
		{
			name: "empty extension",
			sub:  "",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.sub.IsWav()
			if got != tt.want {
				t.Errorf("IsWav(%q) = %v; want %v", tt.sub, got, tt.want)
			}
		})
	}
}

func TestSubtype_IsMpeg(t *testing.T) {
	tests := []struct {
		name string
		sub  Subtype
		want bool
	}{
		{
			name: "valid mpeg",
			sub:  "mpeg",
			want: true,
		},
		{
			name: "valid mp3",
			sub:  "mp3",
			want: true,
		},
		{
			name: "valid mpeg3",
			sub:  "mpeg3",
			want: true,
		},
		{
			name: "valid x-mpeg-3",
			sub:  "x-mpeg-3",
			want: true,
		},
		{
			name: "invalid wav extension",
			sub:  "json",
			want: false,
		},
		{
			name: "empty extension",
			sub:  "",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.sub.IsMpeg()
			if got != tt.want {
				t.Errorf("IsMpeg(%q) = %v; want %v", tt.sub, got, tt.want)
			}
		})
	}
}
