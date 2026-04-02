// Copyright © 2026 Ruslan Sadekov. All rights reserved.

package mime

import (
	"fmt"
	"strings"
)

// ...
const (
	mpegName = "mpeg"
	wavName  = "wav"
	mp4Name  = "mp4"
	oggName  = "ogg"
)

// audioType and videoType are the standard top-level media categories
// defined in RFC 2045/2046.
const (
	audioType = "audio"
	videoType = "video"
)

// Mime represents a full standard media type string (e.g., "audio/mpeg").
type Mime string

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

// Type extracts the primary media category (e.g., "audio", "video") from the MIME.
func (m Mime) Type() (string, error) {
	// Early exit for empty strings to avoid unnecessary processing
	// and provide a clear, actionable error.
	if m == "" {
		return "", ErrEmptyMime
	}

	parts := strings.Split(string(m), "/")
	if len(parts) != 2 {
		return "", ErrInvalidMime
	}

	tp := parts[0]
	if tp == "" {
		return "", ErrEmptyType
	}

	return tp, nil
}

// Subtype extracts the specific format identifier from the MIME.
func (m Mime) Subtype() (Subtype, error) {
	// Early exit for empty strings to avoid unnecessary processing
	// and provide a clear, actionable error.
	if m == "" {
		return "", ErrEmptyMime
	}

	parts := strings.Split(string(m), "/")
	if len(parts) != 2 {
		return "", ErrInvalidMime
	}

	s := parts[1]
	s, _, _ = strings.Cut(s, ";")

	if s == "" {
		return "", ErrEmptySybtype
	}

	return Subtype(s), nil
}

// Parts decomposes a Mime string into its primary type and subtype components.
// It enforces the RFC-standard "type/subtype" format, providing the foundation
// for media categorization and format identification.
func (m Mime) Parts() (string, Subtype, error) {
	// Early exit for empty strings to avoid unnecessary processing
	// and provide a clear, actionable error.
	if m == "" {
		return "", "", ErrEmptyMime
	}

	// We split the string exactly once to maintain high performance.
	// The length check ensures that we have both a category (e.g., "audio")
	// and a specific format (e.g., "mpeg") before proceeding.
	parts := strings.Split(string(m), "/")
	if len(parts) != 2 {
		return "", "", ErrInvalidMime
	}

	tp := parts[0]
	if tp == "" {
		return "", "", ErrEmptyType
	}

	s := parts[1]
	s, _, _ = strings.Cut(s, ";")

	if s == "" {
		return "", "", ErrEmptySybtype
	}

	// Casting the second part to the Subtype custom type enables
	// the use of specialized methods like Subtype.Name().
	return tp, Subtype(s), nil
}

// IsAudio verifies if the MIME type belongs to the "audio" category.
// It returns an error if the MIME structure is malformed, ensuring
// that we distinguish between non-audio files and invalid data.
func (m Mime) IsAudio() (bool, error) {
	t, err := m.Type()
	if err != nil {
		return false, err
	}

	return t == audioType, nil
}

// IsVideo verifies if the MIME type belongs to the "video" category.
// It returns an error if the MIME structure is malformed, ensuring
// that we distinguish between non-video files and invalid data.
func (m Mime) IsVideo() (bool, error) {
	t, err := m.Type()
	if err != nil {
		return false, err
	}

	return t == videoType, nil
}

// MimeIsWav checks if the provided MIME string belongs to the WAV audio family.
// It ensures the media type is 'audio' and the subtype is in the WAV whitelist.
func MimeIsWav(mime Mime) bool {
	if ok, err := mime.IsAudio(); err != nil || !ok {
		return false
	}

	s, err := mime.Subtype()
	if err != nil {
		return false
	}

	// Delegate the actual format check to IsWav for consistent logic reuse.
	return IsWav(Subtype(s))
}

// IsWav checks if the subtype belongs to the WAV family.
func IsWav(s Subtype) bool {
	if s == "" {
		return false
	}

	_, ok := wavSubtypes[s]
	return ok
}

// MimeIsMpeg checks if the provided MIME string belongs to the MPEG audio family.
// It ensures the media type is 'audio' and the subtype is in the MPEG whitelist.
func MimeIsMpeg(mime Mime) bool {
	if ok, err := mime.IsAudio(); err != nil || !ok {
		return false
	}

	s, err := mime.Subtype()
	if err != nil {
		return false
	}

	// Delegate the actual format check to IsMpeg for consistent logic reuse.
	return IsMpeg(Subtype(s))
}

// IsMpeg checks if the subtype belongs to the MPEG family.
func IsMpeg(s Subtype) bool {
	if s == "" {
		return false
	}

	_, ok := mpegSubtypes[s]
	return ok
}
