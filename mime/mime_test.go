// Copyright © 2026 Ruslan Sadekov.

package mime

import "testing"

func TestMime_Type(t *testing.T) {
	tests := []struct {
		name     string
		wantErr  bool
		mime     Mime
		wantType string
	}{
		{
			name:     "valid mime type",
			wantErr:  false,
			mime:     "audio/mpeg",
			wantType: "audio",
		},
		{
			name:     "valid mime type with params",
			wantErr:  false,
			mime:     "audio/mpeg;test=1231",
			wantType: "audio",
		},
		{
			name:    "invalid mime format",
			wantErr: true,
			mime:    "audio/mpeg/invalid",
		},
		{
			name:    "empty mime type",
			wantErr: true,
			mime:    "/mpeg",
		},
		{
			name:    "empty mime",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tp, err := tt.mime.Type()
			if (err != nil) != tt.wantErr {
				t.Errorf("Type(%q): got error %v, want %v", tt.mime, err, tt.wantErr)
				return
			}
			if tt.wantErr {
				return
			}

			if tp != tt.wantType {
				t.Errorf("file type must be: %s, got: %s", tt.wantType, tp)
			}
		})
	}
}

func TestMime_Subtype(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
		mime    Mime
		wantSub Subtype
	}{
		{
			name:    "valid mime type",
			wantErr: false,
			mime:    "audio/mpeg",
			wantSub: "mpeg",
		},
		{
			name:    "valid ogg mime type",
			wantErr: false,
			mime:    "audio/ogg",
			wantSub: "ogg",
		},
		{
			name:    "valid mime type with params",
			wantErr: false,
			mime:    "audio/mpeg;test=1231",
			wantSub: "mpeg",
		},
		{
			name:    "invalid mime format",
			wantErr: true,
			mime:    "audio/mpeg/invalid",
		},
		{
			name:    "empty mime type",
			wantErr: true,
			mime:    "audio/",
		},
		{
			name:    "empty mime",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, err := tt.mime.Subtype()
			if (err != nil) != tt.wantErr {
				t.Errorf("Type(%q): got error %v, want %v", tt.mime, err, tt.wantErr)
				return
			}
			if tt.wantErr {
				return
			}

			if s != tt.wantSub {
				t.Errorf("file type must be: %s, got: %s", tt.wantSub, s)
			}
		})
	}
}

func TestMime_Parts(t *testing.T) {
	tests := []struct {
		name     string
		wantErr  bool
		mime     Mime
		wantType string
		wantSub  Subtype
	}{
		{
			name:     "valid mime type",
			wantErr:  false,
			mime:     "audio/mpeg",
			wantType: "audio",
			wantSub:  "mpeg",
		},
		{
			name:     "valid mime type with params",
			wantErr:  false,
			mime:     "audio/mpeg;test=1231",
			wantType: "audio",
			wantSub:  "mpeg",
		},
		{
			name:    "invalid mime format",
			wantErr: true,
			mime:    "audio/mpeg/invalid",
		},
		{
			name:    "empty mime type",
			wantErr: true,
			mime:    "/mpeg",
		},
		{
			name:    "empty mime subtype",
			wantErr: true,
			mime:    "audio/",
		},
		{
			name:    "empty mime type and subtype",
			wantErr: true,
			mime:    "/",
		},
		{
			name:    "empty mime",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fileType, ext, err := tt.mime.Parts()
			if (err != nil) != tt.wantErr {
				t.Errorf("fileExtension(%q): got error %v, want %v", tt.mime, err, tt.wantErr)
				return
			}
			if tt.wantErr {
				return
			}

			if fileType != tt.wantType {
				t.Errorf("file type must be: %s, got: %s", tt.wantType, fileType)
			}

			if ext != tt.wantSub {
				t.Errorf("file ext must be: %s, got: %s", tt.wantSub, ext)
			}
		})
	}
}

func TestMime_IsAudio(t *testing.T) {
	tests := []struct {
		name string
		mime Mime
		want bool
	}{
		{
			name: "valid",
			mime: "audio/mpeg",
			want: true,
		},
		{
			name: "not audio type",
			mime: "application/mpeg",
			want: false,
		},
		{
			name: "empty type",
			mime: "/mpeg",
			want: false,
		},
		{
			name: "invalid mime format",
			mime: "audio/mpeg/invalid",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ok := tt.mime.IsAudio()

			if tt.want != ok {
				t.Errorf("IsAudio(%q) = %t, want %t", tt.mime, ok, tt.want)
			}
		})
	}
}

func TestMime_IsVideo(t *testing.T) {
	tests := []struct {
		name string
		mime Mime
		want bool
	}{
		{
			name: "valid",
			mime: "video/mpeg",
			want: true,
		},
		{
			name: "not video type",
			mime: "application/mpeg",
			want: false,
		},
		{
			name: "empty type",
			mime: "/mpeg",
			want: false,
		},
		{
			name: "invalid mime format",
			mime: "video/mpeg/invalid",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ok := tt.mime.IsVideo()

			if tt.want != ok {
				t.Errorf("IsVideo(%q) = %t, want %t", tt.mime, ok, tt.want)
			}
		})
	}
}

func TestMime_IsImage(t *testing.T) {
	tests := []struct {
		name string
		mime Mime
		want bool
	}{
		{
			name: "valid",
			mime: "image/png",
			want: true,
		},
		{
			name: "not video type",
			mime: "application/mpeg",
			want: false,
		},
		{
			name: "empty type",
			mime: "/mpeg",
			want: false,
		},
		{
			name: "invalid mime format",
			mime: "video/mpeg/invalid",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ok := tt.mime.IsImage()

			if tt.want != ok {
				t.Errorf("IsImage(%q) = %t, want %t", tt.mime, ok, tt.want)
			}
		})
	}
}
