// Copyright © 2026 Ruslan Sadekov.

package mime

import (
	"errors"
	"strings"

	"vortex.com/files"
)

// TODO:
// ~ IsArchive method
// ~ Subtype param getter (after ';')

const (
	audioType       = "audio"       // audioType represents the base MIME category for sound files.
	videoType       = "video"       // videoType represents the base MIME category for motion picture files.
	imageType       = "image"       // imageType represents the base MIME category for static visual media.
	textType        = "text"        // textType represents the base MIME category for human-readable data.
	applicationType = "application" // applicationType represents binary data or specialized file formats.
	fontType        = "font"        // fontType represents the base MIME category for digital typefaces.
	modelType       = "model"       // modelType represents 3D assets or CAD data.
)

// maxLen defines the upper bound for a MIME string to prevent resource exhaustion.
const maxLen = 1024

// fileTypes maps raw MIME type strings to internal strongly-typed file categories.
var fileTypes = map[string]files.Type{
	audioType: files.Audio,
	videoType: files.Video,
	imageType: files.Image,
	fontType:  files.Font,
}

// Mime represents a full standard media type string (e.g., "audio/mpeg").
type Mime string

// TODO:
// ~ Проверять размер Mime

// Validate ensures the Mime string adheres to the RFC standards (type/subtype).
// It checks for empty values and correctly formatted delimiters, ignoring optional parameters.
// It returns a joined error if either the primary type or the subtype is missing or invalid.
func (m Mime) Validate() error {
	if m == "" {
		return ErrEmptyMime
	}

	if len(m) > maxLen {
		return ErrTooLong
	}

	// A valid MIME type must contain exactly one forward slash separator (type/subtype).
	if count := strings.Count(string(m), "/"); count != 1 {
		return ErrInvalidMime
	}

	t, sub, _ := strings.Cut(string(m), "/")

	// Pre-allocate a slice for potential errors with a capacity of 2.
	// This micro-optimization avoids extra memory allocations.
	var errs = make([]error, 0, 2)

	// The primary type (e.g., "video") must not be an empty string.
	if t == "" {
		errs = append(errs, ErrEmptyType)
	}

	// Strip optional parameters like "; charset=UTF-8" to isolate the subtype.
	sub, _, _ = strings.Cut(sub, ";")
	if sub == "" {
		errs = append(errs, ErrEmptySubtype)
	}

	if len(errs) == 0 {
		return nil
	}
	return errors.Join(errs...)
}

// Subtype extracts the specific format identifier from the MIME.
func (m Mime) Subtype() (Subtype, error) {
	_, sub, err := m.Parts()
	return sub, err
}

// Type extracts the primary media category (e.g., "audio", "video") from the MIME.
func (m Mime) Type() (string, error) {
	t, _, err := m.Parts()
	return t, err
}

// FileType resolves the Mime string into a high-level internal file category.
// It returns files.Unknown and an error if the MIME format is invalid,
// or files.Unknown without an error if the type is valid but not supported by our presets.
func (m Mime) FileType() (files.Type, error) {
	mt, err := m.Type()
	if err != nil {
		// If the MIME string is structurally unsound, we cannot determine its category.
		return files.Unknown, err
	}

	t, found := fileTypes[mt]
	if !found {
		// The MIME is valid (e.g., "application/pdf"), but we don't have a processing category for it.
		return files.Unknown, nil
	}

	return t, nil
}

// Parts decomposes a Mime string into its primary type and subtype components.
// It enforces the RFC-standard "type/subtype" format.
// It returns the primary type, the Subtype, or an error if validation fails.
func (m Mime) Parts() (tp string, sub Subtype, err error) {
	if err := m.Validate(); err != nil {
		return "", "", err
	}

	t, s, _ := strings.Cut(string(m), "/")

	// Strip any optional parameters (like "; charset=utf-8") to isolate the subtype.
	s, _, _ = strings.Cut(s, ";")

	return t, Subtype(s), nil
}

// is checks if the primary media type of the MIME string matches the provided type string.
// It returns false if the MIME format is invalid or if the types do not match.
func (m Mime) is(t string) bool {
	tp, err := m.Type()
	if err != nil {
		return false
	}

	return t == tp
}

// IsAudio returns true if the Mime type is "audio".
func (m Mime) IsAudio() bool { return m.is(audioType) }

// IsVideo returns true if the Mime type is "video".
func (m Mime) IsVideo() bool { return m.is(videoType) }

// IsImage returns true if the Mime type is "image".
func (m Mime) IsImage() bool { return m.is(imageType) }

// IsApplication returns true if the Mime type is "application".
func (m Mime) IsApplication() bool { return m.is(applicationType) }

// IsText returns true if the Mime type is "text".
func (m Mime) IsText() bool { return m.is(textType) }

// IsFont returns true if the Mime type is "font".
func (m Mime) IsFont() bool { return m.is(fontType) }

// IsModel returns true if the Mime type is "model".
func (m Mime) IsModel() bool { return m.is(modelType) }
