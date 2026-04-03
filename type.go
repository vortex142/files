// Copyright © 2026 Ruslan Sadekov.

package files

import (
	"fmt"
)

// TODO:
// ~ Json unmarshal

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
