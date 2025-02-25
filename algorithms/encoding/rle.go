package encoding

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

// Doesn't work with unicode
type RLE struct {
	lengthFirst bool
}

func NewRLE(lengthFirst bool) *RLE {
	return &RLE{
		lengthFirst: lengthFirst,
	}
}

func (rle *RLE) Encode(input []byte) ([]byte, error) {
	// rle works with string, so let's convert to string
	str := string(input)
	var out string
	var r int
	for l := 0; l < len(str); l = r {
		r = l
		for r < len(str) && str[l] == str[r] {
			r++
		}
		if rle.lengthFirst {
			out += strconv.Itoa(r - l)
			out += string(str[l])
		} else {
			out += string(str[l])
			out += strconv.Itoa(r - l)
		}
	}

	return []byte(out), nil
}

func (rle *RLE) Decode(input []byte) ([]byte, error) {
	// rle works with string, so let's convert to string
	str := string(input)
	var out string
	var r int
	for l := 0; l < len(str); l = r {
		r = l
		var num string
		var ch byte
		if !rle.lengthFirst {
			ch = str[r]
			if unicode.IsDigit(rune(ch)) {
				return nil, fmt.Errorf("bad input: first character is a number")
			}
			r++
		}

		for r < len(str) && unicode.IsDigit(rune(str[r])) {
			num += string(str[r])
			r++
		}
		var length int
		if len(num) == 0 {
			length = 1
		} else {
			var err error
			length, err = strconv.Atoi(num)
			if err != nil {
				return nil, err
			}
			if length == 0 {
				return nil, fmt.Errorf("bad input: cannot parse length 0")
			}
		}
		if rle.lengthFirst {
			if length > 1 && r == len(str) { // last charater found was a digit. That's an error
				return nil, fmt.Errorf("bad input: last character is a number")
			}
			ch = str[r]
			r++
		}

		out += strings.Repeat(string(ch), length)
	}

	return []byte(out), nil
}
