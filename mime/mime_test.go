// Copyright © 2026 Ruslan Sadekov.

package mime

import (
	"strings"
	"testing"

	"vortex.com/files"
)

func TestMime_Validate(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		wantErr bool
		mime    Mime
	}{
		{
			name:    "valid mime",
			wantErr: false,
			mime:    "audio/mpeg",
		},
		{
			name:    "valid mime with optional",
			wantErr: false,
			mime:    "video/mp4;test=banana",
		},
		{
			name:    "invalid mime format",
			wantErr: true,
			mime:    "audio/mpeg/invalid",
		},
		{
			name:    "empty mime",
			wantErr: true,
		},
		{
			name:    "empty type",
			wantErr: true,
			mime:    "/mpeg",
		},
		{
			name:    "empty subtype",
			wantErr: true,
			mime:    "audio/",
		},
		{
			name:    "too long mime",
			wantErr: true,
			mime:    Mime(strings.Repeat("audio/mpeg", 500)),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			err := tt.mime.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() = %v, want: %t", err, tt.wantErr)
			}
		})
	}
}

func TestMime_Type(t *testing.T) {
	t.Parallel()

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
			t.Parallel()

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
	t.Parallel()

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
			t.Parallel()

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
	t.Parallel()

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
		{
			name:    "too long mime",
			wantErr: true,
			mime:    Mime(strings.Repeat("audio/mpeg", 500)),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

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

func TestMime_FileType(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		wantErr bool
		mime    Mime
		want    files.Type
	}{
		{
			name:    "empty mime",
			wantErr: true,
		},
		{
			name:    "invalid mime",
			wantErr: true,
			mime:    "audio/mpeg/invalid",
		},
		{
			name:    "empty type (invalid mime)",
			wantErr: true,
			mime:    "/mpeg",
		},
		{
			name:    "empty subtype (invalid mime)",
			wantErr: true,
			mime:    "audio/",
		},
		{
			name:    "audio",
			wantErr: false,
			mime:    "audio/mpeg",
			want:    files.Audio,
		},
		{
			name:    "video",
			wantErr: false,
			mime:    "video/mp4",
			want:    files.Video,
		},
		{
			name:    "image",
			wantErr: false,
			mime:    "image/png",
			want:    files.Image,
		},
		{
			name:    "font",
			wantErr: false,
			mime:    "font/font",
			want:    files.Font,
		},
		{
			name:    "unknown type",
			wantErr: false,
			mime:    "application/zip7",
			want:    files.Unknown,
		},
		{
			name:    "video with optional data",
			wantErr: false,
			mime:    `video/mp4;data="test"`,
			want:    files.Video,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := tt.mime.FileType()
			if (err != nil) != tt.wantErr {
				t.Errorf("FileType() = %v, want: %t", err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}

			if got != tt.want {
				t.Errorf("want type: %s, got: %s", tt.want, got)
			}
		})
	}
}

func TestMime_IsAudio(t *testing.T) {
	t.Parallel()

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
			t.Parallel()

			ok := tt.mime.IsAudio()

			if tt.want != ok {
				t.Errorf("IsAudio(%q) = %t, want %t", tt.mime, ok, tt.want)
			}
		})
	}
}

func TestMime_IsVideo(t *testing.T) {
	t.Parallel()

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
			t.Parallel()

			ok := tt.mime.IsVideo()
			if tt.want != ok {
				t.Errorf("IsVideo(%q) = %t, want %t", tt.mime, ok, tt.want)
			}
		})
	}
}

func TestMime_IsImage(t *testing.T) {
	t.Parallel()

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
			name: "not image type",
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
			t.Parallel()

			ok := tt.mime.IsImage()

			if tt.want != ok {
				t.Errorf("IsImage(%q) = %t, want %t", tt.mime, ok, tt.want)
			}
		})
	}
}

func TestMime_IsApplication(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		mime Mime
		want bool
	}{
		{
			name: "valid",
			mime: "application/zip7",
			want: true,
		},
		{
			name: "not application type",
			mime: "video/mpeg",
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
			t.Parallel()

			ok := tt.mime.IsApplication()

			if tt.want != ok {
				t.Errorf("IsApplication(%q) = %t, want %t", tt.mime, ok, tt.want)
			}
		})
	}
}

func TestMime_IsText(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		mime Mime
		want bool
	}{
		{
			name: "valid",
			mime: "text/txt",
			want: true,
		},
		{
			name: "not text type",
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
			mime: "text/mpeg/invalid",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ok := tt.mime.IsText()

			if tt.want != ok {
				t.Errorf("IsText(%q) = %t, want %t", tt.mime, ok, tt.want)
			}
		})
	}
}

func TestMime_IsFont(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		mime Mime
		want bool
	}{
		{
			name: "valid",
			mime: "font/font",
			want: true,
		},
		{
			name: "not font type",
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
			mime: "font/mpeg/invalid",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ok := tt.mime.IsFont()

			if tt.want != ok {
				t.Errorf("IsFont(%q) = %t, want %t", tt.mime, ok, tt.want)
			}
		})
	}
}

func TestMime_IsModel(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		mime Mime
		want bool
	}{
		{
			name: "valid",
			mime: "model/cad",
			want: true,
		},
		{
			name: "not model type",
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
			mime: "model/mpeg/invalid",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ok := tt.mime.IsModel()
			if tt.want != ok {
				t.Errorf("IsModel(%q) = %t, want %t", tt.mime, ok, tt.want)
			}
		})
	}
}
