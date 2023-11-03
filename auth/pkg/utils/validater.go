package utils

import (
	"fmt"
	"regexp"
)

var (
	isValidUsername    = regexp.MustCompile(`^[a-zA-Z]+$`).MatchString
	isValidPhoneNumber = regexp.MustCompile(`^[0-9]+$`).MatchString
)

func ValidateString(value string, minValue, maxValue int) error {
	if len(value) > maxValue || len(value) < minValue {
		return fmt.Errorf("input string must be min=%d and max=%d", minValue, maxValue)
	}

	return nil
}

func ValidateUsername(username string) error {
	if err := ValidateString(username, 4, 60); err != nil {
		return err
	}

	if !isValidUsername(username) {
		return fmt.Errorf("username must contain only letters")
	}

	return nil
}

func ValidatePassword(password string) error {
	return ValidateString(password, 8, 100)
}

func ValidatePhoneNumber(phoneNumber string) error {
	if err := ValidateString(phoneNumber, 7, 15); err != nil {
		return err
	}

	if !isValidPhoneNumber(phoneNumber) {
		return fmt.Errorf("phone number must contain only digits or spaces")
	}

	return nil
}
