package nhs

import (
	"errors"
	"math/rand"
	"strings"
)

var (
	ErrInvalidLength   = errors.New("invalid length for NHS number")
	ErrExpectedDigit   = errors.New("expected NHS number to only contain digits")
	ErrInvalidChecksum = errors.New("invalid checksum")
)

// ValidateNHSNumber takes an NHS number as a string and returns whether it is a valid number. The string data type is used
// because numbers can start with a zero.
func ValidateNHSNumber(number string) error {
	// Remove spaces to avoid an easy validation issue
	number = strings.ReplaceAll(number, " ", "")

	if len(number) != 10 {
		return ErrInvalidLength
	}

	var checksum int
	var finalDigit int

	for i := 0; i < 10; i++ {
		digitChar := number[i]

		if digitChar < '0' || digitChar > '9' {
			return ErrExpectedDigit
		}

		// Subtract the byte representation of the character '0' to get the decimal value.
		digit := int(digitChar - '0')

		// We don't want to calculate the checksum for the final digit. This check is used, so we can read all the digits
		// in a single loop.
		if i == 9 {
			finalDigit = digit
			break
		}

		// Factor ranges from 10 to 2 (10-0 to 10-8)
		factor := 10 - i
		checksum += digit * factor
	}

	remainder := checksum % 11
	checkDigit := 11 - remainder
	// The specification tells us 0 is used when we get 11.
	if checkDigit == 11 {
		checkDigit = 0
	}

	// Our calculated digit should equal the final number in the NHS number
	if checkDigit != finalDigit {
		return ErrInvalidChecksum
	}

	return nil
}

// GenerateNHSNumber generates an NHS number as a string.
// TODO: This implementation is very naive. I ran out of time so this has the chance to generate invalid NHS numbers.
func GenerateNHSNumber() string {
	var checksum int
	var number string

	for i := 0; i < 9; i++ {
		// Factor ranges from 10 to 2 (10-0 to 10-8)
		factor := 10 - i

		// Generate a number in the range 0 to 9.
		digit := rand.Intn(10)

		checksum += digit * factor

		// Convert to a character by adding the character '0'
		number += string(byte(digit + '0'))
	}

	remainder := checksum % 11
	checkDigit := 11 - remainder
	if checkDigit == 11 {
		checkDigit = 0
	}

	number += string(byte(checkDigit + '0'))

	return number
}
