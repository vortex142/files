// Copyright © 2026 Ruslan Sadekov.

package files

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

// TODO:
// ~ remove regexp in parse

// parsePattern captures the numeric value and the unit suffix.
// It is pre-compiled to mitigate the performance cost during repeated calls.
var parsePattern = regexp.MustCompile(`^([0-9.]+)\s*([a-zA-Z]+)$`)

// nameToUnit maps uppercase unit strings back to their [Unit] constants for fast lookup during parsing.
var nameToUnit = map[string]Unit{
	"B": B, "KB": Kb, "MB": Mb, "GB": Gb, "TB": Tb, "PB": Pb, "EB": Eb,
}

// Unit represents the byte magnitude. It follows the Clean Architecture principle
// of maintaining clear, typed domain boundaries.
//
//go:generate stringer -type=Unit -linecomment
type Unit uint8

const (
	B  Unit = iota // B
	Kb             // KB
	Mb             // MB
	Gb             // GB
	Tb             // TB
	Pb             // PB
	Eb             // EB
)

// Validate checks if the storage unit is defined within the supported enumeration.
// It returns ErrInvalidUnit if the unit exceeds the Eb (Exabyte) threshold.
func (u Unit) Validate() error {
	if u > Eb {
		return ErrInvalidUnit
	}
	return nil
}

// Bytes returns the power-of-two multiplier.
// Values are capped at 2^60 (Exabytes) to prevent uint64 overflow during bit-shifting.
func (u Unit) Bytes() Size {
	if u > 6 {
		u = 6
	}

	// Bitwise shift is used instead of math.Pow for significantly faster O(1) execution
	// at the CPU level.
	return Size(uint64(1) << (10 * u))
}

// Size stores the file size in bytes as a float64, enabling precise fractional.
type Size float64

// String returns a human-readable representation of the Size.
// It uses logarithmic calculation to determine the optimal unit magnitude
// without inefficient iterative loops.
func (s Size) String() string {
	if s == 0 {
		return "0 B"
	}

	// ...
	val := float64(s)
	e := math.Floor(math.Log(val) / math.Log(1024))

	// ...
	if e > 6 {
		e = 6
	}

	// ...
	unit := Unit(e)
	return fmt.Sprintf("%.2f %s", s.To(unit), unit.String())
}

// Validate ensures the Size value resides within a logically sound physical range.
// It returns [ErrNegativeSize] if the value is less than zero, otherwise nil.
func (s Size) Validate() error {
	if s < 0 {
		return ErrNegativeSize
	}
	return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface to support flexible size definitions.
// It allows the Size type to be initialized from either a human-readable string (e.g., "512MB") or a numeric byte value.
// It returns [ErrNegativeSize] if the numeric value is negative, or other errors for invalid formats or data types.
func (s *Size) UnmarshalJSON(data []byte) error {
	var v any
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	switch val := v.(type) {
	case string:
		// Parse human-readable strings like "10GB", "250KB", or "1.5TB" into the internal unit representation.
		return s.Parse(val)
	case float64:
		// Handle raw numeric values, treating them as basic bytes.
		if val < 0 {
			return ErrNegativeSize
		}

		// Normalize the numeric input to the base Byte unit to maintain internal consistency.
		*s = Size(val)
		return nil
	}

	return fmt.Errorf("invalid type for Size: %T", v)
}

// Parse converts a formatted string into a Size value.
// Returns an error if the format or data within the string is invalid.
// NOTE: This is a complex and potentially slow operation because it invokes
// regex matching and string normalization. Avoid calling this within performance-critical
// hot loops to maintain high backend performance.
func (s *Size) Parse(str string) error {
	if str == "" {
		return errors.New("empty size string")
	}

	// Normalization ensures compatibility with nameToUnit lookup regardless of user input style.
	str = strings.ToUpper(strings.TrimSpace(str))

	// The number of matches must be equal to 3, since the first element of the slice is a full string str.
	match := parsePattern.FindStringSubmatch(str)
	if len(match) != 3 {
		return fmt.Errorf(`invalid size format: %s (expected e.g. "1.5 GB")`, str)
	}

	val, err := strconv.ParseFloat(match[1], 64)
	if err != nil {
		return err
	}

	unitStr := match[2]
	unit, ok := nameToUnit[unitStr]
	if !ok {
		return fmt.Errorf("unknown unit: %s", unitStr)
	}

	// Normalizing to bytes provides a single source of truth for all subsequent arithmetic.
	*s = From(Size(val), unit)
	return nil
}

// To converts the current Size (assumed to be in bytes) into the specified target [Unit].
// unit defines the scale (e.g., MB, GB) to which the value should be normalized.
// It returns the scaled Size or 0 if the current value fails validation.
func (s Size) To(unit Unit) float64 {
	// Ensure the size is physically valid (non-negative) before performing division.
	if err := s.Validate(); err != nil {
		return 0
	}

	return float64(s / unit.Bytes())
}

// From creates a Size (in bytes) by scaling a raw value from the provided [Unit].
// val represents the quantity, and unit defines its magnitude (e.g., 5, GB).
// It returns the total byte count or 0 if the input value is invalid.
func From(val Size, unit Unit) Size {
	// Validate the input value to prevent overflow or logic errors from negative numbers.
	if err := val.Validate(); err != nil {
		return 0
	}

	return val * unit.Bytes()
}
