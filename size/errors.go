// Copyright © 2026 Ruslan Sadekov.

package size

import "errors"

// Size error constants define failures related to value validation, unit support, and string parsing for [Size].
var (
	// ErrNegativeSize indicates that the provided file size value is below zero.
	ErrNegativeSize = errors.New("size cannot be negative")

	// ErrInvalidUnit indicates that the provided unit exceeds the maximum supported scale (Exabyte).
	ErrInvalidUnit = errors.New("invalid or unsupported storage unit")

	// ErrEmptyParseStr indicates that the input string is empty and cannot be processed.
	ErrEmptyParseStr = errors.New("parsing empty input string")

	// ErrTooLongParseStr indicates that the input string exceeds the maximum allowed length to prevent resource exhaustion.
	ErrTooLongParseStr = errors.New("parsing string exceeds maximum length")

	// ErrInvalidParseStr indicates that the input string is missing either a numeric value or a unit suffix.
	ErrInvalidParseStr = errors.New("parsing failed due to missing value or unit block")
)
