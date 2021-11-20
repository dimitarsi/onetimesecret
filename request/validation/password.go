package validation

import (
	"strings"
)

const (
	specialChars string = "!@$%^&*()_+~|{}<>?[]'./\\"
)

func CheckPassword(password string) (bool, []string) {
	errors := []string{}
	hasErrors := false

	if len(password) == 0 {
		return true, []string{ "Password is required" }
	} else	if len(password) < 10 {
		hasErrors = true
		errors = append(errors, "Password needs to be at least 10 symbols")
	}

	if strings.ContainsAny(password, specialChars) == false {
		hasErrors = true
		errors = append(errors, "Passwords needs to include at least one special character")
	}

	return hasErrors, errors
}