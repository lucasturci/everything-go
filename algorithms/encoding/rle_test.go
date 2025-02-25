package encoding

import (
	"testing"
)

func TestRLE(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		lengthFirst bool
		wantEncode  string
		wantErr     bool
	}{
		{
			name:        "basic encoding length first",
			input:       "AABBBCCCC",
			lengthFirst: true,
			wantEncode:  "2A3B4C",
			wantErr:     false,
		},
		{
			name:        "basic encoding char first",
			input:       "AABBBCCCC",
			lengthFirst: false,
			wantEncode:  "A2B3C4",
			wantErr:     false,
		},
		{
			name:        "single characters length first",
			input:       "ABC",
			lengthFirst: true,
			wantEncode:  "1A1B1C",
			wantErr:     false,
		},
		{
			name:        "empty string",
			input:       "",
			lengthFirst: true,
			wantEncode:  "",
			wantErr:     false,
		},
		{
			name:        "repeated single char length first",
			input:       "AAAAA",
			lengthFirst: true,
			wantEncode:  "5A",
			wantErr:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rle := NewRLE(tt.lengthFirst)

			// Test encoding
			got, err := rle.Encode([]byte(tt.input))
			if (err != nil) != tt.wantErr {
				t.Errorf("RLE.Encode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if string(got) != tt.wantEncode {
				t.Errorf("RLE.Encode() = %v, want %v", string(got), tt.wantEncode)
			}

			// Test decoding
			decoded, err := rle.Decode(got)
			if err != nil {
				t.Errorf("RLE.Decode() unexpected error = %v", err)
				return
			}
			if string(decoded) != tt.input {
				t.Errorf("RLE.Decode() = %v, want %v", string(decoded), tt.input)
			}
		})
	}
}

func TestRLEDecodeErrors(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		lengthFirst bool
		wantErr     bool
	}{
		{
			name:        "ends with number",
			input:       "A2B3",
			lengthFirst: true,
			wantErr:     true,
		},
		{
			name:        "invalid number",
			input:       "A2B0",
			lengthFirst: false,
			wantErr:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rle := NewRLE(tt.lengthFirst)
			_, err := rle.Decode([]byte(tt.input))
			if (err != nil) != tt.wantErr {
				t.Errorf("RLE.Decode() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
