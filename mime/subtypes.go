// Copyright © 2026 Ruslan Sadekov. All rights reserved.

package mime

// wavSubtypes covers all known variations of the WAV format across
// different OS and browser implementations (RFC 2361 and others).
var wavSubtypes = map[Subtype]struct{}{
	"vnd.wave": {}, "vnd.wav": {}, "wave": {}, "x-wav": {}, "x-wave": {},
}

// mpegSubtypes handles standard and vendor-specific MPEG audio labels.
var mpegSubtypes = map[Subtype]struct{}{
	"mpeg": {}, "mp3": {}, "mpeg3": {}, "x-mpeg-3": {},
}

// oggSubtypes handles standard and vendor-specific OGG audio labels.
var oggSubtypes = map[Subtype]struct{}{
	"ogg": {},
}

// mp4Subtypes handles standard and vendor-specific MP4 video labels.
var mp4Subtypes = map[Subtype]struct{}{
	"mp4": {},
}
