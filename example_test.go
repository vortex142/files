// Copyright © 2026 Ruslan Sadekov.

package files

import (
	"fmt"
)

func ExampleName_CutExtension() {
	name := Name("video.test.mp4")
	fmt.Printf("input string: %q\n", name)

	name.CutExtension()
	fmt.Printf("output string: %q\n", name)

	// Output:
	// input string: "video.test.mp4"
	// output string: "video.test"
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

func ExampleName_Prepare() {
	validName := Name("video.mp4")
	invalidName := Name("../../etc/password")
	reservedName := Name("con")

	oldValidName := validName
	validName.Prepare()
	fmt.Printf("old: %q | prepared: %q\n", oldValidName, validName)

	oldInvalidName := invalidName
	invalidName.Prepare()
	fmt.Printf("old: %q | prepared: %q\n", oldInvalidName, invalidName)

	oldReservedName := reservedName
	reservedName.Prepare()
	fmt.Printf("old: %q | prepared: %q\n", oldReservedName, reservedName)

	// Output:
	// old: "video.mp4" | prepared: "video.mp4"
	// old: "../../etc/password" | prepared: "....etcpassword"
	// old: "con" | prepared: "_con"
}
