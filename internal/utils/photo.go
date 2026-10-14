package utils

import (
	"errors"
	"strings"
)

func ValidateBase64Photo(photo string) error {
	if photo == "" {
		return nil
	}
	parts := strings.Split(photo, ",")
	if len(parts) != 2 {
		return errors.New("invalid base64 photo format")
	}
	return nil
}
