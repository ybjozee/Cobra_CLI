package validation

import (
	"errors"
	"log"
	"regexp"
	"strconv"
)

func GetRotationsFromKey(key string) (int, error) {
	if key == "" {
		return 0, errors.New("required flag \"key\" not provided")
	}
	parsedKey, err := strconv.ParseInt(key, 10, 32)
	if err != nil {
		return 0, errors.New("invalid key provided")
	}
	return int(parsedKey), nil
}

func ValidatePhoneNumber(phoneNumber string) error {
	e164Pattern := `^\+[1-9]\d{1,14}$`
	match, err := regexp.Match(e164Pattern, []byte(phoneNumber))
	if err != nil {
		log.Fatal(err.Error())
	}
	if !match {
		return errors.New("phone number must be in E.164 format")
	}
	return nil
}
