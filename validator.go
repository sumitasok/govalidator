package validator

import (
	"errors"
	"fmt"
	"regexp"
)

type Validate struct {
	ErrorList ErrorList
}

func (v *Validate) MaxLengthOfString(key string, max int, value string) bool {
	key_length := len(value)
	if key_length > max {
		errMsg := fmt.Sprintf("%s: Max allowed length is %d, found %d", key, max, key_length)
		v.ErrorList.Append(key, errors.New(errMsg))
		return false
	}
	return true
}

func (v *Validate) MatchRegExp(key string, exp string, value string) bool {
	errMsg := fmt.Sprintf("%s: Format(%s) doesn't match with (%s)", key, exp, value)
	if matched, err := regexp.MatchString(exp, value); err != nil || matched == false {
		v.ErrorList.Append(key, errors.New(errMsg))
		return false
	}
	return true
}

func (v *Validate) MatchEmail(key string, value string) bool {
	errMsg := fmt.Sprintf("%s: Email format is not correct (%s)", key, value)
	if matched, err := regexp.MatchString("([a-zA-Z0-9])+(@)([a-zA-Z0-9])+((.)[a-zA-Z0-9])+", value); err != nil || matched == false {
		v.ErrorList.Append(key, errors.New(errMsg))
		return false
	}
	return true
}

func (v *Validate) MinLengthOfString(key string, min int, value string) bool {
	key_length := len(value)
	if key_length < min {
		errMsg := fmt.Sprintf("%s: Min required length is %d, found %d", key, min, key_length)
		v.ErrorList.Append(key, errors.New(errMsg))
		return false
	}
	return true
}
