// Copyright © 2026 Ruslan Sadekov.

package files

import (
	"fmt"
	"strings"
)

// TODO:
// ~ Json unmarshal

// typeNames maps string identifiers to their corresponding [Type] constants for unified file classification.
var typeNames = map[string]Type{
	"video":    Video,
	"audio":    Audio,
	"image":    Image,
	"document": Document,
	"archive":  Archive,
	"font":     Font,
}

// The following magic numbers are synchronized with the lengths of identifiers in [typeNames].
// If a new type is added, run tests in type_test.go to verify that these constants
// match the actual minimum and maximum bounds of the map keys.
const (
	// maxTypeLen defines the upper bound for type string validation to prevent processing of abnormally long inputs.
	maxTypeLen = 8

	// minTypeLen sets the minimum expected length for a type identifier to quickly filter out invalid or malformed data.
	minTypeLen = 4
)

// Type represents a high-level classification of a file resource.
// Using uint8 provides a compact representation, ideal for indexing and embedding in larger data structures.
//
//go:generate stringer -type=Type -linecomment
type Type uint8

const (
	Video    Type = iota // video
	Audio                // audio
	Image                // image
	Document             // document
	Archive              // archive
	Font                 // font
	Unknown              // unknown
)

// Validate verifies that the type identifier corresponds to a known file classification.
// It uses the receiver t to perform a range check against the defined constants.
// It returns nil if valid, otherwise [ErrInvalidType] for unrecognized values.
func (t Type) Validate() error {
	if t > Unknown {
		return fmt.Errorf("Validate: %w (received: %d)", ErrInvalidType, t)
	}
	return nil
}

// TypeFromString converts a string identifier into a [Type].
// s represents the raw type name.
// It returns the corresponding [Type] or [Unknown] if no match is found.
//
// Allows any case for searching but requires lowercase input for zero-allocation lookups.
func TypeFromString(s string) Type {
	// Avoids unnecessary map overhead by pre-validating string length against known identifier boundaries.
	if len(s) > maxTypeLen || len(s) < minTypeLen {
		return Unknown
	}

	t, found := typeNames[s]

	// Performs a case-insensitive fallback to ensure compatibility with varied external naming conventions.
	if !found {
		s = strings.ToLower(s)
		t, found = typeNames[s]
	}

	if !found {
		return Unknown
	}
	return t
}
