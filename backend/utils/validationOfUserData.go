package utils

import (
	"errors"
	"regexp"

	"github.com/image-generator/internal/models"
)

// validate user based on his/her name
func validateUser(name string) bool{
	r := regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9]{2,32}$`)

	if r.MatchString(name) {
		return true
	}
	return false
}

// validate the email of the user
func validateEmail(email string) bool{
	r := regexp.MustCompile(`^[a-zA-Z0-9!#$%^&*+.]+@[a-z]+\.(com.np|com|org)$`)

	if r.MatchString(email) {
		return true
	}
	return false
}

// validate if the users password is correct
func validatePassword(password string) bool{
	format := regexp.MustCompile(`^[a-zA-Z0-9!@#$%^&*+.]{8,32}$`)

	var capitalLetters = regexp.MustCompile(`[A-Z]`)
	var smallLetters = regexp.MustCompile(`[a-z]`)
	var numbers = regexp.MustCompile(`[0-9]`)
	var specialCharacters = regexp.MustCompile(`[!@#$%^&*+.]`)


	if format.MatchString(password) && capitalLetters.MatchString(password) && smallLetters.MatchString(password) && numbers.MatchString(password) && specialCharacters.MatchString(password) {
		return true
	}
	return false
}

func ValidateLoginCredentials(loginCredentials models.Login)error{
	if validateEmail(loginCredentials.Email) == false{
		return errors.New("Email must follow the proper format.")
	} else if validatePassword(loginCredentials.Password) == false{
		return errors.New("Password must follow certain format.")
	}
	return nil
}

func ValidateRegisterCredentials(registerCredentials models.Register)error{
	if validateEmail(registerCredentials.Email) == false{
		return errors.New("Email must follow the proper format.")
	} else if validatePassword(registerCredentials.Password) == false{
		return errors.New("Password must follow certain format.")
	} else if validateUser(registerCredentials.Username) == false{
		return errors.New("Username must follow proper format.")
	}
	return nil
}