// Package mime provides tools for parsing, validating, and classifying
// standard MIME.
//
// The package is designed as an extension to [files], allowing for the interpretation
// of file metadata at the web application level. It enables the conversion of string representations
// of types (e.g., "audio/mpeg") into strongly typed [files.Type].
//
// # Key Features
//
//   - Strict validation: Checks for type/subtype conformance according to RFC standards.
//   - Type classification: Matches MIME strings against the internal [files.Type] enumeration.
//   - Deep subtype analysis: Identifies canonical formats (WAV, MPEG) regardless of vendor-specific prefixes.
//   - Security: Built-in string length limit 1024 to prevent resource exhaustion attacks during parsing.
//
// # Performance
//
// The package is optimized for high-load systems. Parsing of components
// is implemented through efficient string splitting without unnecessary memory allocations.
package mime
