package utils

import (
	"errors"
	"net/mail"
	"regexp"
)

type Validator struct{}

func NewValidator() *Validator {
	return &Validator{}
}

func (v *Validator) ValidatePassword(password string) error {
	if len(password) < 8 || len(password) > 32 {
		return errors.New("password must be between 8 and 32 characters")
	}
	hasDigit := regexp.MustCompile("[0-9]").MatchString(password)
	hasCapital := regexp.MustCompile("^[A-Z]").MatchString(password)
	if !hasDigit || !hasCapital {
		return errors.New("password must contain at least one digit and first uppercase letter")
	}
	return nil
}

func (v *Validator) ValidateEmail(email string) error {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return errors.New("invalid email")
	}
	return nil
}
