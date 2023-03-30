package nhs

import "testing"

func TestValidateNHSNumber(t *testing.T) {
	tests := map[string]error{
		"943 476 5919":   nil,
		"5990128088":     nil,
		"1275988113":     nil,
		"4536026665":     nil,
		"7246133720":     nil,
		"5990128087":     ErrInvalidChecksum,
		"4536016660":     ErrInvalidChecksum,
		"invalid number": ErrInvalidLength,
		"987abc4321":     ErrExpectedDigit,
	}

	for input, err := range tests {
		t.Run(input, func(t *testing.T) {
			out := ValidateNHSNumber(input)
			if out != err {
				t.Logf("expected: %v, got: %v", err, out)
				t.Fail()
			}
		})
	}
}
