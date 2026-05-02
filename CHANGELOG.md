# Changelog

## 1.1.1 (2026-05-02)

- optimized `TypeFromString` to avoid unnecessary map lookups
- added `Type.MarshalText` to implement the `encoding.TextMarshaler` interface, allowing types to be serialized as human-readable strings
- added `Type.UnmarshalText` to implement the `encoding.TextUnmarshaler` interface, enabling the use of `Type` as a key in maps
- improved error handling in `UnmarshalJSON` by wrapping errors with `ErrInvalidType` for better debugging and `errors.Is` support

## 1.1.0 (2026-05-02)

- added `Type.UnmarshalJSON` to support polymorphic decoding of types from both strings and numbers
- added `TypeFromString` to parsing human-readable strings like "video", "audio", or "image" into the internal type representation

## v1.0.0 (2026-04-07)

- initial release