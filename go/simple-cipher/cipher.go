package cipher

import (
	"regexp"
	"strings"
)

// Cipher is implemented by simple shift ciphers like Caesar.
// It is expected that Encode will ignore all values in the string that are not
// A-Za-z, they will not be represented in the output. The output will  be also
// normalized to lowercase.
type Cipher interface {
	Encode(string) string
	Decode(string) string
}

// Shift is a more generic shift cipher with a flexible shift distance.
type Shift int

// NewCaesar used to obtain a Shift cipher with a fixed shift distance of 3.
func NewCaesar() Cipher {
	return NewShift(3)
}

// NewShift used to obtain a Shift cipher.
// Argument must be in the range 1 to 25 or -1 to -25, zero is disallowed.
// For invalid arguments returns nil.
func NewShift(distance int) Cipher {
	switch {
	case 1 <= distance && distance <= 25:
		return Shift(distance)
	case -25 <= distance && distance <= -1:
		return Shift(26 + distance)
	default:
		return nil
	}
}

// Encode implements Cipher.Encode.
// All letters will be shifted with the value of Shift.
func (sh Shift) Encode(s string) string {
	return strings.Map(encode(int(sh)), normalize(s))
}

// Decode implements Cipher.Decode. Decodes Shift encoded strings.
func (sh Shift) Decode(s string) string {
	return strings.Map(encode(int(26-sh)), s)
}

// Vigenere is a more complex cipher using a string as key value.
// Each character in the key is used to shift the corresponding character
// by index. If the key is shorter than the text, repeat the key as needed.
type Vigenere string

// NewVigenere used to obtain a Vigenere cipher. The argument  must consist of lower
// case letters a-z only. Values consisting entirely of the letter 'a' are
// disallowed. For invalid arguments it returns nil.
func NewVigenere(key string) Cipher {
	if key == "" || key != normalize(key) || strings.Count(key, "a") == len(key) {
		return nil
	}
	return Vigenere(key)
}

func (v Vigenere) Encode(s string) string {
	var encoded string
	for i, r := range normalize(s) {
		shift := rune(v[i%len(v)] - 'a')
		encoded += string('a' + (r-'a'+shift)%26)
	}
	return encoded
}

func (v Vigenere) Decode(s string) string {
	var decoded string
	for i, r := range normalize(s) {
		shift := rune(v[i%len(v)] - 'a')
		decoded += string('z' - ('z'-r+shift)%26)
	}
	return decoded
}

func normalize(s string) string {
	return strings.ToLower(regexp.MustCompile("[^a-zA-Z]+").ReplaceAllString(s, ""))
}

func encode(n int) func(rune) rune {
	return func(r rune) rune {
		return 'a' + (r-'a'+rune(n))%26
	}
}