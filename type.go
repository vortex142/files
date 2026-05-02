// Copyright © 2026 Ruslan Sadekov.

package files

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

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

// UnmarshalJSON implements the [json.Unmarshaler] interface to support polymorphic decoding of types from both strings and numbers.
// data represents the raw JSON encoded bytes containing either a category name or a numeric identifier.
// It returns an error if the input format is unsupported or if a numeric value is negative.
//
// String values are resolved through [TypeFromString] while numeric values are cast directly and validated.
// Any numeric input that fails validation is safely defaulted to the [Unknown] state.
func (t *Type) UnmarshalJSON(data []byte) error {
	var v any
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	switch val := v.(type) {
	case string:
		*t = TypeFromString(val)
		return nil
	case float64:
		if val < 0 {
			return errors.New("type value cannot be negative")
		}

		*t = Type(val)

		if t.Validate() != nil {
			*t = Unknown
		}

		return nil
	}

	return fmt.Errorf("invalid type for Type: %T", v)
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
		// allLower identifies strings that are already normalized to skip redundant map lookups.
		allLower := true
		for i := 0; i < len(s); i++ {
			if s[i] >= 'A' && s[i] <= 'Z' {
				allLower = false
				break
			}
		}

		// Since all characters are in lowercase and the initial lookup failed,
		// any subsequent search after strings.ToLower will also fail to find a match.
		if allLower {
			return Unknown
		}

		s = strings.ToLower(s)
		t, found = typeNames[s]
	}

	if !found {
		return Unknown
	}
	return t
}
