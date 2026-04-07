# files
[![go tests](https://github.com/vortex142/files/actions/workflows/go-tests.yml/badge.svg)](https://github.com/vortex142/files/actions/workflows/go-tests.yml) [![codecov](https://codecov.io/github/vortex142/files/graph/badge.svg?token=29V742D8P2)](https://codecov.io/github/vortex142/files) [![Go Report Card](https://goreportcard.com/badge/github.com/vortex142/files)](https://goreportcard.com/report/github.com/vortex142/files) [![Doc](https://img.shields.io/badge/godoc-blue.svg)](https://pkg.go.dev/github.com/vortex142/files) ![Go](https://img.shields.io/badge/go-1.24-blue.svg)

[Русский](docs/README.ru.md)

**Files** is a high-performance, I/O-agnostic library for handling file metadata in Go. The package replaces unsafe "raw" string manipulations with a strictly typed domain model, providing built-in name sanitization, unified file type representation, and robust tools for data size management.

The `mime` subpackage extends the core functionality by providing tools for working with MIME strings within the shared domain concept. Integration with the library's internal types allows for direct content classification, replacing redundant linear checks (like `IsAudio` or `IsVideo`) with efficient category mapping.

## Name

The `files` package provides the `Name` type — a smart wrapper over string that handles the "dirty" work of file name validation and normalization, ensuring cross-platform compatibility and security.

### Name Sanitization

The `Sanitize` method strips invalid characters and normalizes the string. If the resulting name matches a system-reserved name (e.g., `con.exe` on Windows), the method safely transforms it into `_con.exe`.

```go
name := Name("../../etc/password")
name.Sanitize() // name -> "....etcpassword"
```

### Name Validation

The `Validate` method checks the name length, presence of prohibited characters, and Path Traversal attempts `(../../)`. It specifically monitors Windows-reserved names (CON, PRN, AUX, etc.), which is critical for servers operating across different OS environments.

```go
name := Name("../../etc/password")
if err := name.Validate(); err != nil {
  // handling invalid file name
}
```

### Removing Extension

The `CutExtension` method extracts and removes the extension, modifying the current name. It correctly handles files with multiple dots, isolating only the final part after the last dot.

```go
n := Name("video.test.mp4")
ext := n.CutExtension() // n -> "video.test", ext = "mp4"
```

## Type

`files.Type` is a uint8 enum that classifies files based on their functional purpose.

### Supported Categories:

- **Audio** — sound files.
- **Video** — video streams and containers.
- **Image** — raster and vector images.
- **Document** — text documents, spreadsheets, presentations.
- **Archive** — compressed packages.
- **Font** — font files.
- **Unknown** — unidentified data.

## `size` Subpackage

The `size` package is designed for working with data volumes in byte representation. Instead of raw integers, it introduces the `Size` type (a wrapper over float64), which handles the routine of conversion, parsing, and formatting.

### `New` Constructor

The `New` constructor allows you to create a `Size` by specifying a value and a unit of measurement (`Unit` — a uint8 type).

```go
s := New(1, Kb) // 1024
s = New(1, Mb) // 1048576
```

### String Parsing

The `FromString` constructor extracts values from strings like "10 MB", "1.5 GB", "GB 10", etc.
> The maximum string length is limited to 1024 bytes to prevent exploitation of the parser with malicious input data.

```go
FromString("10 MB") // s -> 10485760
FromString("1   kb") // s -> 1024
```

### JSON Integration

`Size` implements the `json.Unmarshaler` interface. This allows parsing of configuration files where size can be specified as either a string ("10 GB") or a number (byte count).

```go
type Config struct {
  MaxUploadSize Size `json:"max_upload_size"`
}
```

### Conversion

The `To` method returns the size value in the specified units of measurement.

```go
s := New(1, Mb)
val := s.To(Kb) // val = 1024
```

### Supported Units of Measurement (`Unit`)

| Constant | Value (Bytes) | Description |
| :--- | :--- | :--- |
| `B` | 1 | Byte |
| `Kb` | 1024 | Kilobyte (2¹⁰) |
| `Mb` | 1,048,576 | Megabyte (2²⁰) |
| `Gb` | 1,073,741,824 | Gigabyte (2³⁰) |
| `Tb` | 1,099,511,627,776 | Terabyte (2⁴⁰) |
| `Pb` | 1,125,899,906,842,624 | Petabyte (2⁵⁰) |
| `Eb` | 1,152,921,504,606,846,976 | Exabyte (2⁶⁰) |

## Mime

The `mime` package is an extension of the core `files` package, designed for type-safe handling of MIME strings. The `Mime` type is a wrapper over string.

### Validation

The `Validate` method checks string compliance with RFC standards and verifies the presence of required components (type and subtype).
> Maximum MIME string size is limited to 1024 bytes to prevent parsing exploitation.

```go
m := Mime("audio/mpeg")
if err := m.Validate(); err != nil {
  // handling invalid mime
}
```

### Parsing

The `Type`, `Subtype`, and `Parts` methods allow for efficient separation of a MIME string into its components.

```go
m := Mime("text/plain;charset=utf-8")
t, err := m.Type() // return "text"
s, err := m.Subtype() // return "plain"

// Get all parts in one call
t, s, err = m.Parts()
```

### Category Verification

The `Mime` type provides fast methods for determining content category:

- IsAudio()
- IsVideo()
- IsImage()
- IsText()
- IsApplication()
- IsFont()
- IsModel()

```go
m := Mime("audio/mpeg")
if m.IsAudio() {
  // logic processing audio file
}
```

### `files.Type` Integration

The `FileType` method converts a MIME string into the library's primary enumerable type.

```go
m := Mime("audio/mpeg")
ft := m.FileType() // return files.Audio
```

### Subtype

The package provides the `Subtype` type — a wrapper over `string` used for the identification and classification of media content formats based on their MIME subtypes.

### Name Normalization

Extracts a "canonical" extension or short format name from various MIME subtype variations. This is useful for unifying file names coming from different sources.

```go
wavSub := Subtype("vnd.wave")
mpegSub := Subtype("mpeg3")

n1, err := wavSub.Name() // n1 = "wav"
n2, err := mpegSub.Name() // n2 = "mpeg"
```

### Format Family Verification

The `Subtype` type provides fast semantic methods to determine if content belongs to a specific family, eliminating the need for manual string comparison.

```go
s := Subtype("vnd.wave")
if s.IsWav() {
  // logic processing wav format
}
```

**Available verification methods:**

- IsWav()
- IsMpeg()
- (list to be expanded)

### Extending Support

Registers for existing vendors and their specific subtypes are limited. The package is open for expansion: if you encounter a rare or new subtype, we welcome your Pull Requests adding new mappings to the global registry `(mime/subtype.go)`.

## Errors

### `files` Package

- `ErrReservedName` — returned when a name matches Windows reserved names (e.g., "CON", "NUL").
- `ErrEmptyName` — returned for an empty name string.
- `ErrEmptySanitizedName` — returned if nothing remains after name sanitization.
- `ErrInvalidType` — returned for an unknown file type.

### `size` Package

- `ErrNegativeSize` — returned for negative size values.
- `ErrInvalidUnit` — returned for invalid units of measurement.
- ``ErrEmptyParseStr` — returned for empty parsing strings.
- `ErrTooLongParseStr` — returned when the parsing string exceeds 1024 bytes.
- `ErrInvalidParseStr` — returned when a value or unit is missing in the parsing string.

### `mime` Package

- `ErrEmptyMime` — returned when handling an empty MIME string.
- `ErrInvalidMime` — returned for malformed MIME strings (missing type/subtype).
- `ErrTooLong` — returned when a MIME string exceeds 1024 bytes.
- `ErrEmptyType` — returned for an empty MIME type.
- `ErrEmptySubtype` — returned for an empty MIME subtype.

## Performance

### `files` Package

Name validation results `(Name.Validate)`:

| Scenario | Speed | Memory | Allocations |
| :--- | :--- | :--- | :--- |
| Valid name: "banana-24.05.2022.mp4" | 39.25 ns/op | 0 B/op | 0 allocs/op |
| Reserved name: "con" | 57.25 ns/op | 8 B/op | 1 allocs/op |
| Long name (240 'g' chars)| 637.7 ns/op | 0 B/op | 0 allocs/op |

Name sanitization results `(Name.Sanitize)`:

| Scenario | Speed | Memory | Allocations |
| :--- | :--- | :--- | :--- |
| Valid name: "banana-24.05.2022.mp4" | 40.05 ns/op | 0 B/op | 0 allocs/op |
| Reserved name: "con" | 110.1 ns/op | 32 B/op | 4 allocs/op |
| Name with 240 '#' characters | 479.7 ns/op | 176 B/op | 3 allocs/op |
| Name with 120 '#' and one '1' | 458.0 ns/op | 176 B/op | 3 allocs/op |

### `size` Package

Size parsing results `(Size.FromString)`:

| Scenario | Speed | Memory | Allocations |
| :--- | :--- | :--- | :--- |
| "10 GB" | 38.98 ns/op | 0 B/op | 0 allocs/op |
| "10 gb" | 70.05 ns/op | 8 B/op | 1 allocs/op |
| "10 invalid" | 68.95 ns/op | 8 B/op | 1 allocs/op |
| "1,5 Mb" | 94.49 ns/op | 16 B/op | 2 allocs/op |
> Note: Using lowercase for units increases parsing time.
> Using a comma as a decimal separator increases parsing time as it is replaced with a dot internally.

### `mime` Package

MIME parsing results `(Mime.Parts)`:

Результат парсинга MIME-строки `(Mime.Parts)`:

| Scenario | Speed | Memory | Allocations |
| :--- | :--- | :--- | :--- |
| "audio/mpeg" | 23.48 ns/op | 0 B/op | 0 allocs/op |
| "audio/mpeg;test=banana" | 24.23 ns/op | 0 B/op | 0 allocs/op |
| "/mpeg" | 44.47 ns/op | 40 B/op | 2 allocs/op |

MIME validation results `(Mime.Validate)`:

| Scenario | Speed | Memory | Allocations |
| :--- | :--- | :--- | :--- |
| "audio/mpeg" | 12.36 ns/op | 0 B/op | 0 allocs/op |
| "audio/mpeg;test=banana" | 13.98 ns/op | 0 B/op | 0 allocs/op |
| "/mpeg" | 43.05 ns/op | 40 B/op | 2 allocs/op |

## Links

[Go Doc](https://pkg.go.dev/github.com/vortex142/files)  
[Roadmap](ROADMAP.md)  
[Changelog](CHANGELOG.md)  