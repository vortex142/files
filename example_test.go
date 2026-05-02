// Copyright © 2026 Ruslan Sadekov.

package files

import (
	"encoding/json"
	"fmt"
)

func ExampleName_CutExtension() {
	name := Name("video.test.mp4")

	withoutExt := name
	ext := withoutExt.CutExtension()

	fmt.Printf("input: %q | without ext: %q | ext: %q\n", name, withoutExt, ext)

	// Output:
	// input: "video.test.mp4" | without ext: "video.test" | ext: "mp4"
}

func ExampleName_Validate() {
	validName := Name("video.mp4")
	invalidName := Name("../../etc/password")
	reservedName := Name("con")

	fmt.Printf("input: %q | is valid: %t\n", validName, validName.Validate() == nil)
	fmt.Printf("input: %q | is valid: %t\n", invalidName, invalidName.Validate() == nil)
	fmt.Printf("input: %q | is valid: %t\n", reservedName, reservedName.Validate() == nil)

	// Output:
	// input: "video.mp4" | is valid: true
	// input: "../../etc/password" | is valid: false
	// input: "con" | is valid: false
}

func ExampleName_Sanitize() {
	validName := Name("video.mp4")
	invalidName := Name("../../etc/password")
	reservedName := Name("con")

	oldValidName := validName
	validName.Sanitize()
	fmt.Printf("old: %q | prepared: %q\n", oldValidName, validName)

	oldInvalidName := invalidName
	invalidName.Sanitize()
	fmt.Printf("old: %q | prepared: %q\n", oldInvalidName, invalidName)

	oldReservedName := reservedName
	reservedName.Sanitize()
	fmt.Printf("old: %q | prepared: %q\n", oldReservedName, reservedName)

	// Output:
	// old: "video.mp4" | prepared: "video.mp4"
	// old: "../../etc/password" | prepared: "....etcpassword"
	// old: "con" | prepared: "_con"
}

func ExampleTypeFromString() {
	inputs := []string{
		"video",
		"Video",
		"AUDIO",
		"INVALID",
		"FoNt",
	}

	for _, input := range inputs {
		t := TypeFromString(input)
		fmt.Printf("input: %q | type: %s\n", input, t)
	}

	// Output:
	// input: "video" | type: video
	// input: "Video" | type: video
	// input: "AUDIO" | type: audio
	// input: "INVALID" | type: unknown
	// input: "FoNt" | type: font
}

func ExampleType_UnmarshalJSON() {
	var raw struct {
		T1 Type `json:"type_1"`
		T2 Type `json:"type_2"`
	}

	// e.g. from config
	cfg := `{
		"type_1": "video",
		"type_2": 1
	}` // 0 is audio

	json.Unmarshal([]byte(cfg), &raw)

	fmt.Printf("input: %q | type: %s\ninput: %d | type: %s\n", "video", raw.T1, 1, raw.T2)

	// Output:
	// input: "video" | type: video
	// input: 1 | type: audio
}

func ExampleType_UnmarshalText() {
	var raw struct {
		M map[Type]string `json:"map"`
	}

	// e.g. from config
	cfg := `{
		"map": {
			"video": "I am video ;)",
			"audio": "I am audio :)"
		}
	}`

	json.Unmarshal([]byte(cfg), &raw)

	fmt.Printf("input: %q | type: %s | value: %q\n", "video", "video", raw.M[Video])
	fmt.Printf("input: %q | type: %s | value: %q\n", "audio", "audio", raw.M[Audio])

	// Output:
	// input: "video" | type: video | value: "I am video ;)"
	// input: "audio" | type: audio | value: "I am audio :)"
}

func ExampleType_MarshalText() {
	ts := []Type{Video, Audio, Image, Document, Unknown, 67}

	for _, t := range ts {
		text, err := t.MarshalText()
		if err != nil {
			fmt.Printf("type: %s | error: %v\n", t.String(), err)
			continue
		}

		fmt.Printf("type: %s | text: %q\n", t, text)
	}

	// Output:
	// type: video | text: "video"
	// type: audio | text: "audio"
	// type: image | text: "image"
	// type: document | text: "document"
	// type: unknown | error: unknown file type
	// type: Type(67) | error: unknown file type: (received: 67)
}
