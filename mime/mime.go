// Copyright © 2026 Ruslan Sadekov. All rights reserved.

package mime

import (
	"errors"
	"fmt"
	"strings"
)

// TODO:
// ~ IsArchive method
// ~ Mime param getter

const (
	audioType       = "audio"       // ...
	videoType       = "video"       // ...
	imageType       = "image"       // ...
	textType        = "text"        // ...
	applicationType = "application" // ...
	fontType        = "font"        // ...
	modelType       = "model"       // ...
)

// Mime represents a full standard media type string (e.g., "audio/mpeg").
type Mime string

// ...
func (m Mime) Validate() error {
	var errs = make([]error, 0, 2)

	if m == "" {
		errs = append(errs, ErrEmptyMime)
	}

	// ...
	t, sub, found := strings.Cut(string(m), "/")

	// ...
	if !found {
		errs = append(errs, ErrInvalidMime)
	}

	// ...
	if m != "" && found {
		if t == "" {
			errs = append(errs, ErrEmptyType)
		}

		sub, _, _ = strings.Cut(sub, ";")
		if sub == "" {
			errs = append(errs, ErrEmptySubtype)
		}
	}

	if len(errs) == 0 {
		return nil
	}
	return fmt.Errorf("failed to validate: %w", errors.Join(errs...))
}

// Subtype extracts the specific format identifier from the MIME.
func (m Mime) Subtype() (Subtype, error) {
	// ...
	if err := m.Validate(); err != nil {
		return "", err
	}

	parts := strings.Split(string(m), "/")
	s := parts[1]
	s, _, _ = strings.Cut(s, ";")

	if s == "" {
		return "", ErrEmptySubtype
	}

	return Subtype(s), nil
}

// Type extracts the primary media category (e.g., "audio", "video") from the MIME.
func (m Mime) Type() (string, error) {
	// ...
	if err := m.Validate(); err != nil {
		return "", err
	}

	parts := strings.Split(string(m), "/")
	tp := parts[0]
	if tp == "" {
		return "", ErrEmptyType
	}

	return tp, nil
}

// ...
// func (m Mime) FileType(files.Type, error) {
// }

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
		return "", "", ErrEmptySubtype
	}

	// Casting the second part to the Subtype custom type enables
	// the use of specialized methods like Subtype.Name().
	return tp, Subtype(s), nil
}

// ...
func (m Mime) is(t string) bool {
	tp, err := m.Type()
	if err != nil {
		return false
	}
	return t == tp
}

// ...
func (m Mime) IsAudio() bool { return m.is(audioType) }

// ...
func (m Mime) IsVideo() bool { return m.is(videoType) }

// ...
func (m Mime) IsImage() bool { return m.is(imageType) }

// ...
func (m Mime) IsApplication() bool { return m.is(applicationType) }

// ...
func (m Mime) IsText() bool { return m.is(textType) }

// ...
func (m Mime) IsFont() bool { return m.is(fontType) }

// ...
func (m Mime) IsModel() bool { return m.is(modelType) }
