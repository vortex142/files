// Copyright © 2026 Ruslan Sadekov.

package size

import (
	"encoding/json"
	"fmt"
	"math"
	"strconv"
	"strings"
)

// nameToUnit maps uppercase unit strings back to their [Unit] constants for fast lookup during parsing.
var nameToUnit = map[string]Unit{
	"B": B, "KB": Kb, "MB": Mb, "GB": Gb, "TB": Tb, "PB": Pb, "EB": Eb,
}

// ...
const maxParseLen = 1024

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

// ...
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

// ...
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

// Parse converts a string representation of data amount into a Size value.
// str is a raw string containing a numeric value and a unit suffix.
// it returns an error if the input format is invalid, numeric value is negative or unit is unrecognized.
func (s *Size) Parse(str string) error {
	// Check if the input is empty to avoid unnecessary processing.
	if str == "" {
		return ErrEmptyParseStr
	}

	// Limit string length to prevent resource exhaustion during parsing.
	if len(str) > maxParseLen {
		return ErrTooLongParseStr
	}

	// Initialize indices to isolate numeric and unit parts of the input.
	startNum := -1
	endNum := -1
	blockNum := false

	startLet := -1
	endLet := -1
	blockLet := false

	// Flag to trigger character replacement if localized decimal separators are found.
	needClear := false

	for i := range str {
		c := str[i]

		// Stop scanning if both numeric and unit blocks are already identified.
		if blockLet && blockNum {
			break
		}

		// Limit unit length to two characters to optimize suffix matching.
		if startLet != -1 && (endLet-startLet)+1 >= 2 {
			break
		}

		// Identify numeric characters including '-' and decimal points.
		if ((c >= '0' && c <= '9') || c == '.' || c == ',' || c == '-') && !blockNum {
			// Mark for normalization if a comma is used as a decimal separator.
			if c == ',' {
				needClear = true
			}

			// Track the start and end of the continuous numeric block.
			if startNum == -1 {
				startNum = i
			}

			endNum = i
			continue
		}

		// Mark the end of the first numeric sequence to ensure only one continuous block is parsed.
		blockNum = startNum != -1

		// Identify alphabetic characters for unit identification.
		if ((c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')) && !blockLet {
			// Track the start and end of the continuous unit label block.
			if startLet == -1 {
				startLet = i
			}

			endLet = i
			continue
		}

		// Mark the end of the first letter sequence to prevent parsing multiple disjoint unit suffixes.
		blockLet = startLet != -1
	}

	// Ensure both a value and a unit were found to maintain data integrity.
	if startNum == -1 || startLet == -1 {
		return ErrInvalidParseStr
	}

	numStr := str[startNum : endNum+1]
	unitStr := str[startLet : endLet+1]

	// Replace commas with dots to ensure compatibility with standard float parsers.
	if needClear {
		var b strings.Builder
		b.Grow(len(numStr))

		for i := range numStr {
			c := numStr[i]

			// Normalize decimal separator for strconv.ParseFloat compatibility.
			if c == ',' {
				b.WriteByte('.')
				continue
			}

			b.WriteByte(c)
		}

		numStr = b.String()
	}

	num, err := strconv.ParseFloat(numStr, 64)
	if err != nil {
		return err
	}

	// Validate that the size is non-negative as physical storage cannot be less than zero.
	if num < 0 {
		return ErrNegativeSize
	}

	// Resolve the unit string against supported measurement units using case-insensitive lookup.
	unit, found := nameToUnit[unitStr]
	if !found {
		// Retry with uppercase to support varied naming conventions (like mb or MB).
		unit, found = nameToUnit[strings.ToUpper(unitStr)]
		if !found {
			return ErrInvalidUnit
		}
	}

	*s = From(num, unit)
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
func From(val float64, unit Unit) Size {
	val = max(val, 0)
	return Size(val) * unit.Bytes()
}
