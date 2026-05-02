// Copyright © 2026 Ruslan Sadekov.

package size

import (
	"encoding/json"
	"fmt"
	"math"
	"testing"
)

func TestUnit_String(t *testing.T) {
	var unitNames = map[Unit]string{
		B: "B", Kb: "KB", Mb: "MB", Gb: "GB", Tb: "TB", Pb: "PB", Eb: "EB",
	}

	t.Parallel()

	for i := range uint8(6) {
		unit := Unit(i)
		t.Run(fmt.Sprintf("unit %d", unit), func(t *testing.T) {
			t.Parallel()

			want := unitNames[unit]
			if got := unit.String(); got != want {
				t.Errorf("want string: %s, got: %s", want, got)
			}
		})
	}
}

func TestUnit_Bytes(t *testing.T) {
	t.Parallel()

	for i := range uint(7) {
		pow := float64(i)
		if i > 6 {
			pow = 6
		}
		must := Size(1 * math.Pow(1024, pow))
		unit := Unit(i)

		t.Run(fmt.Sprintf("unit_%s", unit), func(t *testing.T) {
			t.Parallel()

			if unit.Bytes() != must {
				t.Errorf("expected %f bytes for unit %d, got %f", must, i, unit.Bytes())
			}
		})
	}
}

func TestUnit_Validate(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		wantErr bool
		input   Unit
	}{
		{
			name:    "valid b",
			wantErr: false,
			input:   B,
		},
		{
			name:    "valid kb",
			wantErr: false,
			input:   Kb,
		},
		{
			name:    "valid mb",
			wantErr: false,
			input:   Mb,
		},
		{
			name:    "valid gb",
			wantErr: false,
			input:   Gb,
		},
		{
			name:    "valid tb",
			wantErr: false,
			input:   Tb,
		},
		{
			name:    "valid pb",
			wantErr: false,
			input:   Pb,
		},
		{
			name:    "valid eb",
			wantErr: false,
			input:   Eb,
		},
		{
			name:    "invalid unit",
			wantErr: true,
			input:   Eb + 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := tt.input.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() = %v, want: %t", err, tt.wantErr)
			}
		})
	}
}

func TestSize_String(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		size Size
		want string
	}{
		{
			name: "size is zero",
			size: 0,
			want: "0 B",
		},
		{
			name: "1 byte",
			size: 1,
			want: "1.00 B",
		},
		{
			name: "1024 bytes",
			size: 1024,
			want: "1.00 KB",
		},
		{
			name: "2048 bytes",
			size: 2048,
			want: "2.00 KB",
		},
		{
			name: "2048 bytes",
			size: 2048 * 1.25,
			want: "2.50 KB",
		},
		{
			name: "1 megabyte",
			size: New(1, Mb),
			want: "1.00 MB",
		},
		{
			name: "1 gb",
			size: New(1, Gb),
			want: "1.00 GB",
		},
		{
			name: "1 tb",
			size: New(1, Tb),
			want: "1.00 TB",
		},
		{
			name: "1 pb",
			size: New(1, Pb),
			want: "1.00 PB",
		},
		{
			name: "1 eb",
			size: New(1, Eb),
			want: "1.00 EB",
		},
		{
			name: "more 1 eb",
			size: New(1, Eb+1),
			want: "1.00 EB",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := tt.size.String(); got != tt.want {
				t.Errorf("Size(%v).String() = %q; want %q", tt.size, got, tt.want)
			}
		})
	}
}

func TestSize_FromString(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		wantSize Size
		parse    string
		wantErr  bool
	}{
		{
			name:    "empty parse string",
			wantErr: true,
			parse:   "",
		},
		{
			name:    "nagative parse string",
			wantErr: true,
			parse:   "-1.5 MB",
		},
		{
			name:    "invalid value format string",
			wantErr: true,
			parse:   "invalid tb",
		},
		{
			name:    "invalid unit string",
			wantErr: true,
			parse:   "1.5 INVALID",
		},
		{
			name:    "invalid num string",
			wantErr: true,
			parse:   "15-14 Tb",
		},
		{
			name:     "valid num with space",
			wantErr:  false,
			parse:    "15 14 Tb",
			wantSize: New(15, Tb),
		},
		{
			name:     "valid num with ','",
			wantErr:  false,
			parse:    "1,5 Tb",
			wantSize: New(1.5, Tb),
		},
		{
			name:     "string with has > 2 blocks",
			wantErr:  false,
			parse:    "1 tb invalid",
			wantSize: New(1, Tb),
		},
		{
			name:     "valid parse with Cyrillic",
			wantErr:  false,
			wantSize: 1,
			parse:    "1Б B",
		},
		{
			name:     "valid bytes",
			wantErr:  false,
			wantSize: 1,
			parse:    "1 B",
		},
		{
			name:     "valid inverse",
			wantErr:  false,
			wantSize: 1024,
			parse:    "Kb 1",
		},
		{
			name:     "valid bytes with more trims",
			wantErr:  false,
			wantSize: 1,
			parse:    "1      B",
		},
		{
			name:     "valid long bytes",
			wantErr:  false,
			wantSize: 1213123123,
			parse:    "1213123123 B",
		},
		{
			name:     "valid bytes without trims",
			wantErr:  false,
			wantSize: 1,
			parse:    "1B",
		},
		{
			name:     "valid bytes with lower char",
			wantErr:  false,
			wantSize: 1,
			parse:    "1 b",
		},
		{
			name:     "valid kb",
			wantErr:  false,
			wantSize: New(1, Kb),
			parse:    "1 KB",
		},
		{
			name:     "valid fractional kb",
			wantErr:  false,
			wantSize: New(1.77, Kb),
			parse:    "1.77 KB",
		},
		{
			name:     "valid mb",
			wantErr:  false,
			wantSize: New(1, Mb),
			parse:    "1 MB",
		},
		{
			name:     "valid gb",
			wantErr:  false,
			wantSize: New(1, Gb),
			parse:    "1 GB",
		},
		{
			name:     "valid tb",
			wantErr:  false,
			wantSize: New(1, Tb),
			parse:    "1 TB",
		},
		{
			name:     "valid pb",
			wantErr:  false,
			wantSize: New(1, Pb),
			parse:    "1 PB",
		},
		{
			name:     "valid eb",
			wantErr:  false,
			wantSize: New(1, Eb),
			parse:    "1 EB",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			s, err := FromString(tt.parse)

			if (err != nil) != tt.wantErr {
				t.Errorf("Parse(%s) = %v; wantErr = %t", tt.parse, err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}

			if s != tt.wantSize {
				t.Errorf("Parse(%s) = %v; want = %v", tt.parse, s, tt.wantSize)
			}
		})
	}
}

func TestSize_New(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		want Size
		val  float64
		unit Unit
	}{
		{
			name: "negative",
			want: 0,
			val:  -10000,
			unit: B,
		},
		{
			name: "value is zero",
			want: 0,
			val:  0,
			unit: B,
		},
		{
			name: "1 byte",
			want: 1,
			val:  1,
			unit: B,
		},
		{
			name: "1024 bytes",
			want: 1024,
			val:  1024,
			unit: B,
		},
		{
			name: "1 kb",
			want: 1024,
			val:  1,
			unit: Kb,
		},
		{
			name: "1.5 kb",
			want: 1024 * 1.5,
			val:  1.5,
			unit: Kb,
		},
		{
			name: "1 mb",
			want: Size(math.Pow(1024, 2)),
			val:  1,
			unit: Mb,
		},
		{
			name: "1 gb",
			want: Size(math.Pow(1024, 3)),
			val:  1,
			unit: Gb,
		},
		{
			name: "1 tb",
			want: Size(math.Pow(1024, 4)),
			val:  1,
			unit: Tb,
		},
		{
			name: "1 pb",
			want: Size(math.Pow(1024, 5)),
			val:  1,
			unit: Pb,
		},
		{
			name: "1 eb",
			want: Size(math.Pow(1024, 6)),
			val:  1,
			unit: Eb,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			size := New(tt.val, tt.unit)
			if size != tt.want {
				t.Errorf("From(%v, %s) = %v; want %v", tt.val, tt.unit, size, tt.want)
			}
		})
	}
}

func TestSize_To(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		size Size
		want float64
		unit Unit
	}{
		{
			name: "size is negative",
			size: -100,
			want: 0,
			unit: B,
		},
		{
			name: "size is zero",
			size: 0,
			want: 0,
			unit: B,
		},
		{
			name: "1 b to b",
			size: 1,
			want: 1,
			unit: B,
		},
		{
			name: "1024b to b",
			size: 1024,
			want: 1024,
			unit: B,
		},
		{
			name: "1024b to kb",
			size: 1024,
			want: 1,
			unit: Kb,
		},
		{
			name: "1024*1.5b to kb",
			size: 1024 * 1.5,
			want: 1.5,
			unit: Kb,
		},
		{
			name: "1mb from bytes to mb",
			size: Size(math.Pow(1024, 2)),
			want: 1,
			unit: Mb,
		},
		{
			name: "1.5mb from bytes to mb",
			size: Size(math.Pow(1024, 2)) * 1.5,
			want: 1.5,
			unit: Mb,
		},
		{
			name: "1gb from bytes to gb",
			size: Size(math.Pow(1024, 3)),
			want: 1,
			unit: Gb,
		},
		{
			name: "tb from bytes to tb",
			size: Size(math.Pow(1024, 4)),
			want: 1,
			unit: Tb,
		},
		{
			name: "1pb from bytes to pb",
			size: Size(math.Pow(1024, 5)),
			want: 1,
			unit: Pb,
		},
		{
			name: "1eb from bytes to eb",
			size: Size(math.Pow(1024, 6)),
			want: 1,
			unit: Eb,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			val := tt.size.To(tt.unit)
			if val != tt.want {
				t.Errorf("To(%s) = %f, got: %f", tt.unit, val, tt.want)
			}
		})
	}
}

func TestSize_Validate(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		wantErr bool
		input   Size
	}{
		{
			name:    "zero size",
			wantErr: false,
			input:   0,
		},
		{
			name:    "positive size",
			wantErr: false,
			input:   131231232,
		},
		{
			name:    "negative size",
			wantErr: true,
			input:   -1000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			err := tt.input.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() = %v, want: %t", err, tt.wantErr)
			}
		})
	}
}

func TestSize_UnmarshalJson(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		wantErr bool
		input   any
		want    Size
	}{
		{
			name:    "valid string size (1mb)",
			wantErr: false,
			input:   "1 MB",
			want:    New(1, Mb),
		},
		{
			name:    "valid string size (1.5gb)",
			wantErr: false,
			input:   "1.5 GB",
			want:    New(1.5, Gb),
		},
		{
			name:    "empty string",
			wantErr: true,
			input:   "",
		},
		{
			name:    "invalid string",
			wantErr: true,
			input:   "INVALID",
		},
		{
			name:    "negative string (-1mb)",
			wantErr: true,
			input:   "-1 MB",
		},
		{
			name:    "valid number",
			wantErr: false,
			input:   100,
			want:    100,
		},
		{
			name:    "negative number",
			wantErr: true,
			input:   -100,
		},
		{
			name:    "zero number",
			wantErr: false,
			input:   0,
			want:    0,
		},
		{
			name:    "invalid type",
			wantErr: true,
			input:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			data, err := json.Marshal(tt.input)
			if err != nil {
				t.Fatalf("failed to marshal input: %v", err)
			}

			got := Size(0)
			err = json.Unmarshal(data, &got)

			if (err != nil) != tt.wantErr {
				t.Errorf("Unmarshal() = %v, want: %t", err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}

			if got != tt.want {
				t.Errorf("want: %v, got: %v", tt.want, got)
			}
		})
	}
}
