// Package files provides high-performance utilities for handling file metadata.
//
// Unlike the standard 'os' package, this library is strictly filesystem-agnostic.
// It focuses on the logical representation of files — their names and types —
// without performing any actual I/O operations.
//
// # Core Components
//
// The package is built around two main domain types:
//
//   - Name: Handles secure, cross-platform file naming, preventing path traversal
//     and ensuring compatibility with Windows reserved names.
//     for human-readable parsing and automated unit scaling.
//   - Type: Offers a compact, enumerated classification system for file resources.
//
// # Performance
//
// Performance is a primary goal. The package avoids unnecessary allocations
// by utilizing hand-rolled scanners, bit-shifting for arithmetic, and compact
// primitive types (uint8) for domain enumerations.
package files
