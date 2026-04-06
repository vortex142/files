// Copyright © 2026 Ruslan Sadekov.

package size

import (
	"encoding/json"
	"fmt"
)

func ExampleNew() {
	val := float64(10) // e.g. from config
	unit := Mb

	s := New(val, unit)
	fmt.Printf("input: %0.0f %s | size: %0.0f B\n", val, unit, s)

	// Output:
	// input: 10 MB | size: 10485760 B
}

func ExampleFromString() {
	// e.g. from configs
	strs := []string{"100 gb", "GB 100", "15,6mb", "14.2   MB"}

	for _, str := range strs {
		s, err := FromString(str)
		if err != nil {
			// handling the parsing error
		}

		fmt.Printf("input: %q | size: %0.0f\n", str, s)
	}

	// Output:
	// input: "100 gb" | size: 107374182400
	// input: "GB 100" | size: 107374182400
	// input: "15,6mb" | size: 16357786
	// input: "14.2   MB" | size: 14889779
}

func ExampleSize_To() {
	val := float64(1) // e.g. from config
	u := Mb
	s := New(val, u)
	to := s.To(Kb)

	fmt.Printf("input: %0.0f %s | size: %0.0f B | to: %0.0f %s", val, u, s, to, Kb)

	// Output:
	// input: 1 MB | size: 1048576 B | to: 1024 KB
}

func ExampleSize_String() {
	s1 := New(1, Kb)
	s2 := New(15, Mb)

	fmt.Printf("size: %0.0f | string: %s\nsize: %0.0f | string: %s\n", s1, s1, s2, s2)

	// Output:
	// size: 1024 | string: 1.00 KB
	// size: 15728640 | string: 15.00 MB
}

func ExampleSize_Validate() {
	s1 := Size(-100)
	s2 := Size(100)

	fmt.Printf("size: %0.0f | valid: %t\nsize: %0.0f | valid: %t\n", s1, s1.Validate() == nil, s2, s2.Validate() == nil)

	// Output:
	// size: -100 | valid: false
	// size: 100 | valid: true
}

func ExampleSize_UnmarshalJSON() {
	var raw struct {
		S1 Size `json:"size_1"`
		S2 Size `json:"size_2"`
	}

	// e.g. from config
	cfg := `{
		"size_1": "10 Gb",
		"size_2": 1231312
	}`

	json.Unmarshal([]byte(cfg), &raw)

	fmt.Printf("input: %q | size: %0.0f\ninput: %0.0f | size: %0.0f\n", "10 Gb", raw.S1, float64(1231312), raw.S2)

	// Output:
	// input: "10 Gb" | size: 10737418240
	// input: 1231312 | size: 1231312
}

func ExampleUnit_Bytes() {
	u1 := B
	u2 := Kb
	u3 := Mb

	fmt.Printf("input: %q | bytes: %0.0f\n", u1, u1.Bytes())
	fmt.Printf("input: %q | bytes: %0.0f\n", u2, u2.Bytes())
	fmt.Printf("input: %q | bytes: %0.0f\n", u3, u3.Bytes())

	// Output:
	// input: "B" | bytes: 1
	// input: "KB" | bytes: 1024
	// input: "MB" | bytes: 1048576
}
