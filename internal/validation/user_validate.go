package validation

import (
	"errors"
)

func ValidatePassword(password string) error {
	if len(password) < 6 {
		return errors.New("❌ Password must be at least 6 characters")
	}
	return nil
}
