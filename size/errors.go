// Copyright © 2026 Ruslan Sadekov.

package size

import "errors"

// ...
var (
	// ErrNegativeSize indicates that the provided file size value is below zero.
	ErrNegativeSize = errors.New("size cannot be negative")

	// ErrInvalidUnit indicates that the provided unit exceeds the maximum supported scale (Exabyte).
	ErrInvalidUnit = errors.New("invalid or unsupported storage unit")
)
