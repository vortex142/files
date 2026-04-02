// Copyright © 2026 Ruslan Sadekov. All rights reserved.

package mime

import "errors"

var (
	// ErrEmptyMime is returned when the input string is empty.
	ErrEmptyMime = errors.New("mime is required")

	// ErrInvalidMime is returned when the string exists but doesn't follow the 'type/subtype' structure (e.g., missing the forward slash).
	ErrInvalidMime = errors.New("mime type must be in the format 'type/subtype'")

	// ErrEmptyType is returned when the output MIME type is empty.
	ErrEmptyType = errors.New("mime type is empty")

	// ErrEmptySubtype is returned when the output MIME subtype is empty.
	ErrEmptySubtype = errors.New("mime subtype is empty")
)
