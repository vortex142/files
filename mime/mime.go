// Copyright © 2026 Ruslan Sadekov. All rights reserved.

package mime

import (
	"errors"
	"fmt"
	"strings"
)

// TODO:
// ~ IsArchive method
// ~ Subtype param getter (after ';')

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
	// ...
	if m == "" {
		return ErrEmptyMime
	}

	// ...
	t, sub, found := strings.Cut(string(m), "/")

	// ...
	if !found {
		return ErrInvalidMime
	}

	var errs = make([]error, 0, 2)

	// ...
	if found {
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
	_, sub, err := m.Parts()
	return sub, err
}

// Type extracts the primary media category (e.g., "audio", "video") from the MIME.
func (m Mime) Type() (string, error) {
	t, _, err := m.Parts()
	return t, err
}

// ...
// func (m Mime) FileType(files.Type, error) {
// }

// Parts decomposes a Mime string into its primary type and subtype components.
// It enforces the RFC-standard "type/subtype" format, providing the foundation
// for media categorization and format identification.
func (m Mime) Parts() (tp string, sub Subtype, err error) {
	// ...
	if err := m.Validate(); err != nil {
		return "", "", err
	}

	// ...
	t, s, _ := strings.Cut(string(m), "/")

	// ...
	s, _, _ = strings.Cut(s, ";")

	return t, Subtype(s), nil
	// // the use of specialized methods like Subtype.Name().
	// return tp, Subtype(s), nil
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
