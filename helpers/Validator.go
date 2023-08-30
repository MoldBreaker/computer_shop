package helpers

import (
	"errors"
	"golang.org/x/exp/slices"
	"mime/multipart"
	"regexp"
	"strconv"
	"strings"
)

type Validator struct {
	Chain []error
}

func (Validator *Validator) Required(value string, returnStr ...string) error {
	if len(value) == 0 {
		if returnStr[0] == "" {
			return errors.New("'" + value + "' is empty")
		} else {
			return errors.New(returnStr[0])
		}
	}
	return nil
}

func (Validator *Validator) IsEmail(email string, returnStr ...string) error {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	if !emailRegex.MatchString(email) {
		if returnStr[0] == "" {
			return errors.New("'" + email + "' is not a valid email")
		} else {
			return errors.New(returnStr[0])
		}
	} else {
		return nil
	}
}

func (Validator *Validator) MinLength(value string, min int, returnStr ...string) error {
	if len(value) < min {
		if returnStr[0] == "" {
			return errors.New("'" + value + "' need at least " + strconv.Itoa(min) + " characters")
		} else {
			return errors.New(returnStr[0])
		}
	}
	return nil
}

func (Validator *Validator) MaxLength(value string, max int, returnStr ...string) error {
	if len(value) > max {
		if returnStr[0] == "" {
			return errors.New("'" + value + "' not more than " + strconv.Itoa(max) + " characters")
		} else {
			return errors.New(returnStr[0])
		}
	}
	return nil
}

func (Validator *Validator) ComfirmPassword(password, comfirm_password string, returnStr ...string) error {
	if password != comfirm_password {
		if returnStr[0] == "" {
			return errors.New("password not match")
		} else {
			return errors.New(returnStr[0])
		}
	}
	return nil
}

func (Validator *Validator) IsImage(file *multipart.FileHeader) error {
	ext := strings.Split(file.Filename, ".")
	commonExt := []string{"jpeg", "png", "jpg"}
	if !slices.Contains(commonExt, ext[len(ext)-1]) {
		return errors.New("'" + file.Filename + "' is not an image")
	}
	return nil
}

func (Validator *Validator) Validate() error {
	for i := 0; i < len(Validator.Chain); i++ {
		err := Validator.Chain[i]
		if err != nil {
			return err
		}
	}
	return nil
}
