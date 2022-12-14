package validator

import (
	"errors"
	"github.com/aesoper101/go-utils/regexpx"
	"github.com/aesoper101/go-utils/str"
	"github.com/asaskevich/govalidator"
)

func PhoneValidator(phone string) error {
	if str.IsNotEmpty(phone) && !regexpx.IsPhoneNumber(phone) {
		return errors.New("invalid phone number")
	}
	return nil
}

func UrlValidator(url string) error {
	if str.IsNotEmpty(url) && !govalidator.IsURL(url) {
		return errors.New("invalid url address")
	}
	return nil
}

func IPValidator(ip string) error {
	if str.IsNotEmpty(ip) && !govalidator.IsIP(ip) {
		return errors.New("invalid ip address")
	}
	return nil
}

func EmailValidator(email string) error {
	if str.IsNotEmpty(email) && !regexpx.IsEmail(email) {
		return errors.New("invalid email address")
	}
	return nil
}
