// Copyright © 2026 Ruslan Sadekov. All rights reserved.

package files

import (
	"errors"
	"fmt"
)

// ErrInvalidType indicates that the provided file classification is not recognized by the system.
var ErrInvalidType = errors.New("unknown file type")

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
)

// Validate verifies that the type identifier corresponds to a known file classification.
// It uses the receiver t to perform a range check against the defined constants.
// It returns nil if valid, otherwise ErrInvalidType for unrecognized values.
func (t Type) Validate() error {
	if t > Font {
		return fmt.Errorf("Validate: %w (received: %d)", ErrInvalidType, t)
	}
	return nil
}
