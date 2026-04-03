// Package files provides high-performance utilities for handling file metadata.
//
// Unlike the standard 'os' package, this library is strictly filesystem-agnostic.
// It focuses on the logical representation of files — their names, sizes, and types —
// without performing any actual I/O operations.
//
// # Core Components
//
// The package is built around three main domain types:
//
//   - Name: Handles secure, cross-platform file naming, preventing path traversal
//     and ensuring compatibility with Windows reserved names.
//   - Size: Provides a type-safe way to represent and convert file sizes
//     with support for human-readable parsing (e.g., "1.5 GB").
//   - Type: Offers a compact, enumerated classification system for file resources.
//
// # Performance
//
// Performance is a primary goal. The package avoids unnecessary allocations
// by using [strings.Builder], bit-shifting for unit conversion, and compact
// primitive types (uint8) for enumerations.
package files
