// Copyright © 2026 Ruslan Sadekov.

package files

import "errors"

// These errors cover naming validation, type integrity, and physical constraints.
var (
	// ErrReservedName is returned when the input matches OS-protected strings (e.g., "CON", "NUL").
	ErrReservedName = errors.New("name is reserved")

	// ErrEmptyName is returned when the input string is empty before any processing.
	ErrEmptyName = errors.New("name is required")

	// ErrEmptyPreparedName indicates that the sanitization process stripped all characters, resulting in an unusable identifier.
	ErrEmptyPreparedName = errors.New("prepared name is empty")

	// ErrInvalidType indicates that the provided file classification is not recognized by the system.
	ErrInvalidType = errors.New("unknown file type")
)
