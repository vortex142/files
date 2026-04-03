// Copyright © 2026 Ruslan Sadekov.
// Contributions are welcome! If you find a missing format, please open a PR.

package mime

import (
	"fmt"
)

// wavSubtypes covers all known variations of the WAV format across
// different OS and browser implementations (RFC 2361 and others).
var wavSubtypes = map[Subtype]struct{}{
	"vnd.wave": {}, "vnd.wav": {}, "wave": {}, "x-wav": {}, "x-wave": {},
}

// mpegSubtypes handles standard and vendor-specific MPEG audio labels.
var mpegSubtypes = map[Subtype]struct{}{
	"mpeg": {}, "mp3": {}, "mpeg3": {}, "x-mpeg-3": {},
}

// oggSubtypes handles standard and vendor-specific OGG audio labels.
var oggSubtypes = map[Subtype]struct{}{
	"ogg": {},
}

// mp4Subtypes handles standard and vendor-specific MP4 video labels.
var mp4Subtypes = map[Subtype]struct{}{
	"mp4": {},
}

// ...
const (
	mpegName = "mpeg"
	wavName  = "wav"
	mp4Name  = "mp4"
	oggName  = "ogg"
)

// Subtype represents the second part of a MIME type (e.g., "x-wav" in "audio/x-wav").
// Identifying the subtype is key to determining the correct file extension.
type Subtype string

// Name maps the specific MIME subtype to a canonical format name.
// Unlike a simple getter, it returns an error for unrecognized subtypes,
// allowing the caller to decide whether to use a fallback or reject the file.
func (s Subtype) Name() (string, error) {
	// Level 1: Check for WAV-family formats.
	if _, ok := wavSubtypes[s]; ok {
		return wavName, nil
	}

	// Level 2: Check for MPEG-family formats.
	if _, ok := mpegSubtypes[s]; ok {
		return mpegName, nil
	}

	// Level 3: check from OGG-family formats.
	if _, ok := oggSubtypes[s]; ok {
		return oggName, nil
	}

	// Level 4: check from MP4-family formats.
	if _, ok := mp4Subtypes[s]; ok {
		return mp4Name, nil
	}

	// Instead of returning a "magic string" like "unknown", we return
	// a formatted error. This makes the system's limitations
	// discoverable and easier to debug.
	return "", fmt.Errorf("unknown subtype: %s", s)
}

// IsMpeg checks if the subtype belongs to the Wav family.
func (s Subtype) IsWav() bool {
	if s == "" {
		return false
	}

	_, ok := wavSubtypes[s]
	return ok
}

// IsMpeg checks if the subtype belongs to the MPEG family.
func (s Subtype) IsMpeg() bool {
	if s == "" {
		return false
	}

	_, ok := mpegSubtypes[s]
	return ok
}
