package helpers

import (
	"errors"
	"regexp"
	"strings"
)

type ValidationResult map[string]string

func ValidateFields(fields map[string]string) (string, error) {
	errorsResults := ValidationResult{}

	for field, value := range fields {
		var error string
		switch field {
		case "email":
			error = ValidateEmail(value)
		case "password":
			error = ValidatePassword(value)
		case "username":
			error = ValidateUsername(value)
		case "name", "firstName", "lastName":
			error = ValidateName(value)
		default:
			continue
		}
		if error != "" {
			errorsResults[field] = error
		}
	}

	if len(errorsResults) > 0 {
		return errorsResults.ToString(), errors.New("validation failed")
	}

	return "", nil
}

func (vr ValidationResult) ToString() string {
	var sb strings.Builder
	for field, err := range vr {
		sb.WriteString(field + ": " + err + "\n")
	}
	return sb.String()
}

func ValidateEmail(email string) string {
	if email == "" {
		return "email is required"
	}

	if len(email) > 255 {
		return "email is too long"
	}

	emailRegex := regexp.MustCompile(`^[^\s@]+@[^\s@]+\.[^\s@]+$`)
	if !emailRegex.MatchString(email) {
		return "email is invalid"
	}

	return ""
}

func ValidatePassword(password string) string {
	if password == "" {
		return "password is required"
	}

	if len(password) < 6 {
		return "password is too short"
	}

	if len(password) > 255 {
		return "password is too long"
	}

	uppercaseRegex := regexp.MustCompile(`[A-Z]`)
	if !uppercaseRegex.MatchString(password) {
		return "password must contain at least one uppercase letter"
	}

	lowercaseRegex := regexp.MustCompile(`[a-z]`)
	if !lowercaseRegex.MatchString(password) {
		return "password must contain at least one lowercase letter"
	}

	numberRegex := regexp.MustCompile(`[0-9]`)
	if !numberRegex.MatchString(password) {
		return "password must contain at least one number"
	}

	specialCharacterRegex := regexp.MustCompile(`[^A-Za-z0-9]`)
	if !specialCharacterRegex.MatchString(password) {
		return "password must contain at least one special character"
	}

	return ""
}

func ValidateName(name string) string {
	if name == "" {
		return "Name is required"
	}

	if len(name) > 255 {
		return "Name is too long"
	}

	if len(name) < 3 {
		return "Name should be at least 3 characters long"
	}

	numberRegex := regexp.MustCompile(`[0-9]`)
	if numberRegex.MatchString(name) {
		return "Name must not contain numbers"
	}

	specialCharacterRegex := regexp.MustCompile(`[^A-Za-z\s]`)
	if specialCharacterRegex.MatchString(name) {
		return "Name must not contain special characters"
	}

	return ""
}

func ValidateUsername(username string) string {
	if username == "" {
		return "Username is required"
	}

	if len(username) > 255 {
		return "Username is too long"
	}

	if len(username) < 3 {
		return "Username is too short"
	}

	usernameRegex := regexp.MustCompile(`^[a-z0-9_]+$`)
	if !usernameRegex.MatchString(username) {
		return "Username must contain only lowercase letters, numbers, and underscores"
	}

	return ""
}
