// Copyright © 2026 Ruslan Sadekov.

package mime

import "fmt"

func ExampleMime_Validate() {
	// Simulating multiple MIME type checks
	inputs := []Mime{
		"audio/mpeg",
		"video/mp4",
		"text/plain;charset=utf-8",
		"invalid/mime/test/go",
		"",
	}

	for _, in := range inputs {
		err := in.Validate()
		if err != nil {
			fmt.Printf("Input: %q | Status: Failed\n", in)
			continue
		}

		fmt.Printf("Input: %q | Status: Valid\n", in)
	}

	// Output:
	// Input: "audio/mpeg" | Status: Valid
	// Input: "video/mp4" | Status: Valid
	// Input: "text/plain;charset=utf-8" | Status: Valid
	// Input: "invalid/mime/test/go" | Status: Failed
	// Input: "" | Status: Failed
}

func ExampleMime_Parts() {
	// Simulating multiple MIME type checks
	inputs := []Mime{
		"audio/mpeg",
		"video/mp4",
		"text/plain;charset=utf-8",
		"invalid/mime/test/go",
		"",
	}

	for _, in := range inputs {
		t, sub, err := in.Parts()
		if err != nil {
			fmt.Printf("Input: %q | Status: Failed\n", in)
			continue
		}

		fmt.Printf("Input: %q | type: %s | subtype: %s | Status: Valid\n", in, t, sub)
	}

	// Output:
	// Input: "audio/mpeg" | type: audio | subtype: mpeg | Status: Valid
	// Input: "video/mp4" | type: video | subtype: mp4 | Status: Valid
	// Input: "text/plain;charset=utf-8" | type: text | subtype: plain | Status: Valid
	// Input: "invalid/mime/test/go" | Status: Failed
	// Input: "" | Status: Failed
}

func ExampleMime_FileType() {
	// Simulating multiple MIME type checks
	inputs := []Mime{
		"audio/mpeg",
		"video/mp4",
		"font/ttf",
		"text/plain;charset=utf-8",
		"invalid/mime/test/go",
		"",
	}

	for _, in := range inputs {
		t, err := in.FileType()
		if err != nil {
			fmt.Printf("Input: %q | Status: Failed\n", in)
			continue
		}

		fmt.Printf("Input: %q | type: %s | Status: Valid\n", in, t)
	}

	// Output:
	// Input: "audio/mpeg" | type: audio | Status: Valid
	// Input: "video/mp4" | type: video | Status: Valid
	// Input: "font/ttf" | type: font | Status: Valid
	// Input: "text/plain;charset=utf-8" | type: unknown | Status: Valid
	// Input: "invalid/mime/test/go" | Status: Failed
	// Input: "" | Status: Failed
}

func ExampleMime_IsAudio() {
	// Simulating multiple MIME type checks
	inputs := []Mime{
		"audio/mpeg",
		"video/mp4",
	}

	for _, in := range inputs {
		fmt.Printf("Input: %q | Status: %t\n", in, in.IsAudio())
	}

	// Output:
	// Input: "audio/mpeg" | Status: true
	// Input: "video/mp4" | Status: false
}

func ExampleMime_IsVideo() {
	// Simulating multiple MIME type checks
	inputs := []Mime{
		"audio/mpeg",
		"video/mp4",
	}

	for _, in := range inputs {
		fmt.Printf("Input: %q | Status: %t\n", in, in.IsVideo())
	}

	// Output:
	// Input: "audio/mpeg" | Status: false
	// Input: "video/mp4" | Status: true
}

func ExampleMime_IsImage() {
	// Simulating multiple MIME type checks
	inputs := []Mime{
		"audio/mpeg",
		"image/png",
	}

	for _, in := range inputs {
		fmt.Printf("Input: %q | Status: %t\n", in, in.IsImage())
	}

	// Output:
	// Input: "audio/mpeg" | Status: false
	// Input: "image/png" | Status: true
}

func ExampleMime_IsApplication() {
	// Simulating multiple MIME type checks
	inputs := []Mime{
		"audio/mpeg",
		"application/7zip",
	}

	for _, in := range inputs {
		fmt.Printf("Input: %q | Status: %t\n", in, in.IsApplication())
	}

	// Output:
	// Input: "audio/mpeg" | Status: false
	// Input: "application/7zip" | Status: true
}

func ExampleMime_IsText() {
	// Simulating multiple MIME type checks
	inputs := []Mime{
		"audio/mpeg",
		"text/plain",
	}

	for _, in := range inputs {
		fmt.Printf("Input: %q | Status: %t\n", in, in.IsText())
	}

	// Output:
	// Input: "audio/mpeg" | Status: false
	// Input: "text/plain" | Status: true
}

func ExampleMime_IsFont() {
	// Simulating multiple MIME type checks
	inputs := []Mime{
		"audio/mpeg",
		"font/ttf",
	}

	for _, in := range inputs {
		fmt.Printf("Input: %q | Status: %t\n", in, in.IsFont())
	}

	// Output:
	// Input: "audio/mpeg" | Status: false
	// Input: "font/ttf" | Status: true
}

func ExampleMime_IsModel() {
	// Simulating multiple MIME type checks
	inputs := []Mime{
		"audio/mpeg",
		"model/3mf",
	}

	for _, in := range inputs {
		fmt.Printf("Input: %q | Status: %t\n", in, in.IsModel())
	}

	// Output:
	// Input: "audio/mpeg" | Status: false
	// Input: "model/3mf" | Status: true
}

func ExampleSubtype_IsMpeg() {
	// Simulating multiple MIME type checks
	inputs := []Mime{
		"audio/mpeg",
		"audio/x-mpeg-3",
		"video/mp4",
		"audio/mpeg/invalid",
	}

	for _, in := range inputs {
		sub, err := in.Subtype()
		if err != nil {
			fmt.Printf("Input: %q | Status: Failed\n", in)
			continue
		}
		fmt.Printf("Input: %q | Status: %t\n", in, sub.IsMpeg())
	}

	// Output:
	// Input: "audio/mpeg" | Status: true
	// Input: "audio/x-mpeg-3" | Status: true
	// Input: "video/mp4" | Status: false
	// Input: "audio/mpeg/invalid" | Status: Failed
}

func ExampleSubtype_IsWav() {
	// Simulating multiple MIME type checks
	inputs := []Mime{
		"audio/vnd.wave",
		"audio/wave",
		"audio/x-wav",
		"audio/ogg",
		"audio/mpeg/invalid",
	}

	for _, in := range inputs {
		sub, err := in.Subtype()
		if err != nil {
			fmt.Printf("Input: %q | Status: Failed\n", in)
			continue
		}
		fmt.Printf("Input: %q | Status: %t\n", in, sub.IsWav())
	}

	// Output:
	// Input: "audio/vnd.wave" | Status: true
	// Input: "audio/wave" | Status: true
	// Input: "audio/x-wav" | Status: true
	// Input: "audio/ogg" | Status: false
	// Input: "audio/mpeg/invalid" | Status: Failed
}

func ExampleSubtype_Name() {
	// Simulating multiple MIME type checks
	inputs := []Mime{
		"audio/vnd.wave",
		"audio/wave",
		"audio/x-wav",
		"audio/ogg",
		"audio/mpeg",
		"audio/mp3",
		"audio/mpeg3",
		"audio/mpeg/invalid",
	}

	for _, in := range inputs {
		sub, err := in.Subtype()
		if err != nil {
			fmt.Printf("Input: %q | Status: Failed\n", in)
			continue
		}

		n, err := sub.Name()
		if err != nil {
			fmt.Printf("Input: %q | Status: Failed\n", in)
			continue
		}

		fmt.Printf("Input: %q | Status: %s\n", in, n)
	}

	// Output:
	// Input: "audio/vnd.wave" | Status: wav
	// Input: "audio/wave" | Status: wav
	// Input: "audio/x-wav" | Status: wav
	// Input: "audio/ogg" | Status: ogg
	// Input: "audio/mpeg" | Status: mpeg
	// Input: "audio/mp3" | Status: mpeg
	// Input: "audio/mpeg3" | Status: mpeg
	// Input: "audio/mpeg/invalid" | Status: Failed
}
