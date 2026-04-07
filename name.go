// Copyright © 2026 Ruslan Sadekov.

package files

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

// forbiddenWindowsNames contains legacy MS-DOS device names that are reserved
// by the Windows API. Blocking these is essential for cross-platform stability.
var forbiddenWindowsNames = map[string]struct{}{
	"CON": {}, "PRN": {}, "AUX": {}, "NUL": {},
	"COM1": {}, "COM2": {}, "COM3": {}, "COM4": {}, "COM5": {}, "COM6": {}, "COM7": {}, "COM8": {}, "COM9": {},
	"LPT1": {}, "LPT2": {}, "LPT3": {}, "LPT4": {}, "LPT5": {}, "LPT6": {}, "LPT7": {}, "LPT8": {}, "LPT9": {},
}

// MaxNameLen is set to 240 to ensure compatibility with most filesystems (limit 255).
// We reserve 15 characters for extensions and temporary suffixes added during processing.
const MaxNameLen = 240

// prepareBufferScale defines the multiplier for initial memory allocation
// in strings.Builder. This specific overhead is a deliberate trade-off to
// eliminate the risk of additional heap allocations during string sanitization.
const prepareBufferScale = 2

// Name represents a technical filename derived from metadata or system defaults.
type Name string

// Validate performs a multi-layer security and compatibility check on the filename.
// It ensures the name is not empty, adheres to filesystem length limits, and
// contains no illegal characters that could lead to injection or system errors.
func (n Name) Validate() error {
	// Immediate check for existence to prevent downstream logic from
	// processing empty identifiers.
	if n == "" {
		return ErrEmptyName
	}

	// Cross-platform guard: ensures the name doesn't clash with Windows
	// reserved device names, which is critical for distributed worker nodes.
	if isReservedWinName(n) {
		return ErrReservedName
	}

	// Optimization: we first check the byte length because a UTF-8 rune
	// is always at least 1 byte. If len(n) is within MaxNameLen,
	// RuneCount is guaranteed to be within limits as well, allowing us
	// to skip the expensive string traversal.
	if len(n) > MaxNameLen {
		// Only if the byte count exceeds the limit do we perform
		// a precise count of multi-byte characters to ensure valid UTF-8 length.
		if count := utf8.RuneCountInString(string(n)); count > MaxNameLen {
			return fmt.Errorf("name is too long: %d (max %d)", count, MaxNameLen)
		}
	}

	// Whitelist filtering prevents path traversal and potential shell
	// command injections in the media processing pipeline.
	for _, r := range n {
		if !isAllowed(r) {
			return fmt.Errorf("contains invalid character: %c", r)
		}
	}

	return nil
}

// Prepare sanitizes the filename in-place, transforming potentially dangerous
// or incompatible input into a safe format. Unlike Validate, this method
// proactively "fixes" the name to maintain a smooth user experience.
func (n *Name) Prepare() error {
	// Quick check to avoid unnecessary work.
	if err := n.Validate(); err == nil {
		return nil
	}

	v := *n

	// Initial fallback: if the name is missing, we immediately assign the
	// provided default to avoid working with an empty receiver.
	if v == "" {
		return ErrEmptyName
	}

	var b strings.Builder
	// Pre-allocating memory with a scale factor ensures that even if certain
	// characters are expanded or replaced during sanitization, the builder
	// won't trigger additional, expensive heap allocations inside the loop.
	b.Grow(min(len(v), MaxNameLen*prepareBufferScale))

	count := 0
	for _, r := range v {
		// Hard limit on character count prevents DoS attacks via oversized
		// strings and ensures compatibility with Telegram's UI constraints.
		if count >= MaxNameLen {
			break
		}

		// Line breaks are neutralized to prevent formatting issues in
		// log files and unexpected behavior in the Telegram message display.
		if r == '\n' || r == '\r' {
			b.WriteRune('_')
			count++
			continue
		}

		// Only whitelisted characters are preserved to ensure the resulting
		// string is safe for cross-platform filesystem paths.
		if isAllowed(r) {
			b.WriteRune(r)
			count++
		}
	}

	// If sanitization removes all characters, we return a fallback to
	// avoid empty identifiers that could break database logic.
	if b.Len() == 0 {
		return ErrEmptyPreparedName
	}

	res := Name(b.String())

	// Windows reserved names (like NUL or CON) are escaped with a prefix.
	// This ensures that media can be processed on Windows-based worker nodes
	// without triggering OS-level access violations.
	if isReservedWinName(res) {
		*n = "_" + res
		return nil
	}

	*n = res
	return nil
}

// CutExtension modifies the [Name] by removing its extension and returning it as an [Extension] type.
// n is a pointer to the filename string that will be truncated in place.
// it returns an [Extension] containing the characters after the last dot or an empty string if no dot exists.
func (n *Name) CutExtension() Extension {
	s := string(*n)

	// We search for the last occurrence of a dot to avoid
	// over-trimming names that contain versioning or dates.
	if i := strings.LastIndex(s, "."); i != -1 {
		// Slice the string up to the last dot, effectively removing
		// the extension while preserving the semantic filename.
		*n = Name(s[:i])
		return Extension(s[i+1:])
	}

	return ""
}

// isReservedWinName identifies legacy system filenames that are prohibited
// by the Windows API. This check is vital for cross-platform stability.
func isReservedWinName(n Name) bool {
	// Quick length filter to avoid expensive map lookups for strings
	// that cannot possibly be reserved names.
	if len(n) < 3 || len(n) > 4 {
		return false
	}

	// Case-insensitive comparison ensures that "nul", "NUL", and "Nul"
	// are all correctly identified as dangerous.
	_, ok := forbiddenWindowsNames[strings.ToUpper(string(n))]
	return ok
}

// isAllowed defines the character security policy for the file.
// A layered approach is used: first, high-speed ASCII filtering, then
// broader Unicode support for handling international file names.
func isAllowed(r rune) bool {
	// Level 1: Fast-path for standard Latin alphanumeric characters.
	// This covers the majority of technical names with zero overhead.
	if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') {
		return true
	}

	// Level 2: Whitelist of safe punctuation marks.
	// These are essential for file naming conventions and basic UI formatting
	// while remaining safe from path traversal or shell injection.
	if r == '.' || r == '_' || r == '-' || r == ' ' || r == '(' || r == ')' {
		return true
	}

	// Level 3: International Support.
	// We allow any valid Unicode letters (e.g., Cyrillic) or numbers.
	// We remove potentially dangerous control characters or emoji.
	return unicode.IsLetter(r) || unicode.IsDigit(r)
}
