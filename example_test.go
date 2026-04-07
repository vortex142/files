// Copyright © 2026 Ruslan Sadekov.

package files

import (
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
