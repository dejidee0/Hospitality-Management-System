package utils

import "regexp"

// ValidateEmail checks if the given email is valid.
func ValidateEmail(email string) bool {
	// Regular expression for validating email
	const emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	// Compile the regex
	re := regexp.MustCompile(emailRegex)

	// Match the email against the regex
	return re.MatchString(email)
}
