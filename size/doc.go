// Package size provides comprehensive tools for parsing, and converting
// digital storage magnitudes (file sizes).
//
// The package focuses on a clear domain representation of data amounts,
// abstracting the complexities of unit scales and string manipulation.
//
// # Core Components
//
// The functionality is centered around two main types:
//   - Unit: An enumeration representing magnitudes from Bytes to Exabytes.
//   - Size: A numeric type representing a specific amount of data in bytes.
//
// # String Parsing
//
// The package offers robust parsing of human-readable strings into machine-usable
// byte counts. Key features of the [FromString] function include:
//   - Format Flexibility: Supports both spaced ("1.5 GB") and joined ("100MB") inputs.
//   - Internationalization: Automatically handles both dots (".") and commas (",")
//     as decimal separators (e.g., "1,5 TB" is equivalent to "1.5 TB").
//   - Resilience: Employs case-insensitive unit lookup, correctly resolving
//     "kb", "KB", and "Kb" to the same magnitude.
//   - Validation: Rejects invalid formats, negative values, and unrecognized
//     unit suffixes.
//
// # Data Conversion
//
// Converting between different measurement scales is handled through a type-safe
// API:
//   - Scaling Up: The [New] function creates a [Size] by scaling a raw value
//     by a given [Unit] (e.g., creating a 5 GB representation).
//   - Scaling Down: The [To] method normalizes a [Size] into a specific
//     target unit for reporting or logic.
//
// # Human-Readable Output
//
// The [Size.String] method automatically selects the appropriate unit of measurement based on the amount of data (for example, 1024 bytes is converted to "1.00 KB").
//
// # JSON Integration
//
// To simplify API development, the [Size] type natively supports JSON
// (un)marshaling. It can be initialized from either a numeric byte value
// or a human-readable string directly during JSON decoding.
package size
