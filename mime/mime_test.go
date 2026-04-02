// // Copyright © 2026 Ruslan Sadekov. All rights reserved.

package mime

// import "testing"

// func TestSubtype_Name(t *testing.T) {
// 	tests := []struct {
// 		name    string
// 		subtype Subtype
// 		want    string
// 		wantErr bool
// 	}{
// 		{"vnd.wave", "vnd.wave", wavName, false},
// 		{"x-wav", "x-wav", wavName, false},
// 		{"mpeg", "mpeg", mpegName, false},
// 		{"mp3", "mp3", mpegName, false},
// 		{"mp4", "mp4", videoType, false},
// 		{"ogg", "ogg", oggName, false},
// 		{"unknown", "unknown", "", true},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got, err := tt.subtype.Name()
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("Name(%s) = %s, err: %v, wantErr %t", tt.subtype, got, err, tt.wantErr)
// 			}
// 		})
// 	}
// }

// func TestMime_Type(t *testing.T) {
// 	tests := []struct {
// 		name     string
// 		wantErr  bool
// 		mime     Mime
// 		wantType string
// 	}{
// 		{
// 			name:     "valid mime type",
// 			wantErr:  false,
// 			mime:     "audio/mpeg",
// 			wantType: "audio",
// 		},
// 		{
// 			name:     "valid mime type with params",
// 			wantErr:  false,
// 			mime:     "audio/mpeg;test=1231",
// 			wantType: "audio",
// 		},
// 		{
// 			name:    "invalid mime format",
// 			wantErr: true,
// 			mime:    "audio/mpeg/invalid",
// 		},
// 		{
// 			name:    "empty mime type",
// 			wantErr: true,
// 			mime:    "/mpeg",
// 		},
// 		{
// 			name:    "empty mime",
// 			wantErr: true,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			tp, err := tt.mime.Type()
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("Type(%q): got error %v, want %v", tt.mime, err, tt.wantErr)
// 				return
// 			}
// 			if tt.wantErr {
// 				return
// 			}

// 			if tp != tt.wantType {
// 				t.Errorf("file type must be: %s, got: %s", tt.wantType, tp)
// 			}
// 		})
// 	}
// }

// func TestMime_Subtype(t *testing.T) {
// 	tests := []struct {
// 		name    string
// 		wantErr bool
// 		mime    Mime
// 		wantSub Subtype
// 	}{
// 		{
// 			name:    "valid mime type",
// 			wantErr: false,
// 			mime:    "audio/mpeg",
// 			wantSub: "mpeg",
// 		},
// 		{
// 			name:    "valid ogg mime type",
// 			wantErr: false,
// 			mime:    "audio/ogg",
// 			wantSub: "ogg",
// 		},
// 		{
// 			name:    "valid mime type with params",
// 			wantErr: false,
// 			mime:    "audio/mpeg;test=1231",
// 			wantSub: "mpeg",
// 		},
// 		{
// 			name:    "invalid mime format",
// 			wantErr: true,
// 			mime:    "audio/mpeg/invalid",
// 		},
// 		{
// 			name:    "empty mime type",
// 			wantErr: true,
// 			mime:    "audio/",
// 		},
// 		{
// 			name:    "empty mime",
// 			wantErr: true,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			s, err := tt.mime.Subtype()
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("Type(%q): got error %v, want %v", tt.mime, err, tt.wantErr)
// 				return
// 			}
// 			if tt.wantErr {
// 				return
// 			}

// 			if s != tt.wantSub {
// 				t.Errorf("file type must be: %s, got: %s", tt.wantSub, s)
// 			}
// 		})
// 	}
// }

// func TestMime_Parts(t *testing.T) {
// 	tests := []struct {
// 		name     string
// 		wantErr  bool
// 		mime     Mime
// 		wantType string
// 		wantSub  Subtype
// 	}{
// 		{
// 			name:     "valid mime type",
// 			wantErr:  false,
// 			mime:     "audio/mpeg",
// 			wantType: "audio",
// 			wantSub:  "mpeg",
// 		},
// 		{
// 			name:     "valid mime type with params",
// 			wantErr:  false,
// 			mime:     "audio/mpeg;test=1231",
// 			wantType: "audio",
// 			wantSub:  "mpeg",
// 		},
// 		{
// 			name:    "invalid mime format",
// 			wantErr: true,
// 			mime:    "audio/mpeg/invalid",
// 		},
// 		{
// 			name:    "empty mime type",
// 			wantErr: true,
// 			mime:    "/mpeg",
// 		},
// 		{
// 			name:    "empty mime subtype",
// 			wantErr: true,
// 			mime:    "audio/",
// 		},
// 		{
// 			name:    "empty mime type and subtype",
// 			wantErr: true,
// 			mime:    "/",
// 		},
// 		{
// 			name:    "empty mime",
// 			wantErr: true,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			fileType, ext, err := tt.mime.Parts()
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("fileExtension(%q): got error %v, want %v", tt.mime, err, tt.wantErr)
// 				return
// 			}
// 			if tt.wantErr {
// 				return
// 			}

// 			if fileType != tt.wantType {
// 				t.Errorf("file type must be: %s, got: %s", tt.wantType, fileType)
// 			}

// 			if ext != tt.wantSub {
// 				t.Errorf("file ext must be: %s, got: %s", tt.wantSub, ext)
// 			}
// 		})
// 	}
// }

// func TestMime_IsWav(t *testing.T) {
// 	tests := []struct {
// 		name string
// 		mime Mime
// 		want bool
// 	}{
// 		{
// 			name: "valid audio/vnd.wave",
// 			mime: "audio/vnd.wave",
// 			want: true,
// 		},
// 		{
// 			name: "valid audio/vnd.wav",
// 			mime: "audio/vnd.wav",
// 			want: true,
// 		},
// 		{
// 			name: "valid audio/wave",
// 			mime: "audio/wave",
// 			want: true,
// 		},
// 		{
// 			name: "valid audio/x-wav",
// 			mime: "audio/x-wav",
// 			want: true,
// 		},
// 		{
// 			name: "valid audio/x-wave",
// 			mime: "audio/x-wave",
// 			want: true,
// 		},
// 		{
// 			name: "valid audio/x-wave with charset=UTF-8",
// 			mime: "audio/x-wave;charset=UTF-8",
// 			want: true,
// 		},
// 		{
// 			name: "invalid mime type",
// 			mime: "application/json",
// 			want: false,
// 		},
// 		{
// 			name: "empty mime type",
// 			mime: "",
// 			want: false,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got := MimeIsWav(tt.mime)
// 			if got != tt.want {
// 				t.Errorf("MimeIsWav(%q) = %v; want %v", tt.mime, got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestMime_IsMpeg(t *testing.T) {
// 	tests := []struct {
// 		name string
// 		mime Mime
// 		want bool
// 	}{
// 		{
// 			name: "valid audio/mpeg",
// 			mime: "audio/mpeg",
// 			want: true,
// 		},
// 		{
// 			name: "valid audio/mp3",
// 			mime: "audio/mp3",
// 			want: true,
// 		},
// 		{
// 			name: "valid audio/mpeg3",
// 			mime: "audio/mpeg3",
// 			want: true,
// 		},
// 		{
// 			name: "valid audio/x-mpeg-3",
// 			mime: "audio/x-mpeg-3",
// 			want: true,
// 		},
// 		{
// 			name: "valid audio/mpeg with charset=UTF-8",
// 			mime: "audio/mpeg;charset=UTF-8",
// 			want: true,
// 		},
// 		{
// 			name: "invalid mime type",
// 			mime: "application/json",
// 			want: false,
// 		},
// 		{
// 			name: "empty mime type",
// 			mime: "",
// 			want: false,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got := MimeIsMpeg(tt.mime)
// 			if got != tt.want {
// 				t.Errorf("MimeIsMpeg(%q) = %v; want %v", tt.mime, got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestMime_IsAudio(t *testing.T) {
// 	tests := []struct {
// 		name    string
// 		mime    Mime
// 		want    bool
// 		wantErr bool
// 	}{
// 		{
// 			name:    "valid",
// 			mime:    "audio/mpeg",
// 			want:    true,
// 			wantErr: false,
// 		},
// 		{
// 			name:    "not audio type",
// 			mime:    "application/mpeg",
// 			want:    false,
// 			wantErr: false,
// 		},
// 		{
// 			name:    "empty type",
// 			mime:    "/mpeg",
// 			want:    false,
// 			wantErr: true,
// 		},
// 		{
// 			name:    "invalid mime format",
// 			mime:    "audio/mpeg/invalid",
// 			want:    true,
// 			wantErr: true,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			ok, err := tt.mime.IsAudio()
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("IsAudio(%q) = %t, wantErr %t", tt.mime, ok, tt.wantErr)
// 			}
// 			if tt.wantErr {
// 				return
// 			}

// 			if tt.want != ok {
// 				t.Errorf("IsAudio(%q) = %t, want %t", tt.mime, ok, tt.want)
// 			}
// 		})
// 	}
// }

// func TestMime_IsVideo(t *testing.T) {
// 	tests := []struct {
// 		name    string
// 		mime    Mime
// 		want    bool
// 		wantErr bool
// 	}{
// 		{
// 			name:    "valid",
// 			mime:    "video/mpeg",
// 			want:    true,
// 			wantErr: false,
// 		},
// 		{
// 			name:    "not video type",
// 			mime:    "application/mpeg",
// 			want:    false,
// 			wantErr: false,
// 		},
// 		{
// 			name:    "empty type",
// 			mime:    "/mpeg",
// 			want:    false,
// 			wantErr: true,
// 		},
// 		{
// 			name:    "invalid mime format",
// 			mime:    "video/mpeg/invalid",
// 			want:    true,
// 			wantErr: true,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			ok, err := tt.mime.IsVideo()
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("IsVideo(%q) = %t, wantErr %t", tt.mime, ok, tt.wantErr)
// 			}
// 			if tt.wantErr {
// 				return
// 			}

// 			if tt.want != ok {
// 				t.Errorf("IsVideo(%q) = %t, want %t", tt.mime, ok, tt.want)
// 			}
// 		})
// 	}
// }

// func TestSubtype_IsWav(t *testing.T) {
// 	tests := []struct {
// 		name string
// 		sub  Subtype
// 		want bool
// 	}{
// 		{
// 			name: "valid vnd.wave",
// 			sub:  "vnd.wave",
// 			want: true,
// 		},
// 		{
// 			name: "valid vnd.wav",
// 			sub:  "vnd.wav",
// 			want: true,
// 		},
// 		{
// 			name: "valid wave",
// 			sub:  "wave",
// 			want: true,
// 		},
// 		{
// 			name: "valid x-wav",
// 			sub:  "x-wav",
// 			want: true,
// 		},
// 		{
// 			name: "valid x-wave",
// 			sub:  "x-wave",
// 			want: true,
// 		},
// 		{
// 			name: "invalid wav extension",
// 			sub:  "json",
// 			want: false,
// 		},
// 		{
// 			name: "empty extension",
// 			sub:  "",
// 			want: false,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got := IsWav(tt.sub)
// 			if got != tt.want {
// 				t.Errorf("IsWav(%q) = %v; want %v", tt.sub, got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestSubtype_IsMpeg(t *testing.T) {
// 	tests := []struct {
// 		name string
// 		sub  Subtype
// 		want bool
// 	}{
// 		{
// 			name: "valid mpeg",
// 			sub:  "mpeg",
// 			want: true,
// 		},
// 		{
// 			name: "valid mp3",
// 			sub:  "mp3",
// 			want: true,
// 		},
// 		{
// 			name: "valid mpeg3",
// 			sub:  "mpeg3",
// 			want: true,
// 		},
// 		{
// 			name: "valid x-mpeg-3",
// 			sub:  "x-mpeg-3",
// 			want: true,
// 		},
// 		{
// 			name: "invalid wav extension",
// 			sub:  "json",
// 			want: false,
// 		},
// 		{
// 			name: "empty extension",
// 			sub:  "",
// 			want: false,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got := IsMpeg(tt.sub)
// 			if got != tt.want {
// 				t.Errorf("IsMpeg(%q) = %v; want %v", tt.sub, got, tt.want)
// 			}
// 		})
// 	}
// }
