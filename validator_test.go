package validator

import (
	// "fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStruct(t *testing.T) {
	assert := assert.New(t)

	validate := Validate{}
	assert.NotNil(validate)

	maxLen := validate.MaxLengthOfString("key", 10, "valuevaluevaluevalue")
	assert.False(maxLen)
}

func TestCountPresence(t *testing.T) {
	assert := assert.New(t)

	validate := Validate{}
	assert.NotNil(validate)

	element := 1
	list := []int{1, 2, 3}
	actual := validate.CountPresence("key", element, list)
	assert.Equal(true, actual)

	validate = Validate{}
	assert.NotNil(validate)

	element_str := "sumit"
	list_str := []string{"sumit", "tiger"}
	actual = validate.CountPresence("key", element_str, list_str)
	assert.Equal(true, actual)

	// var element_interface interface{} = "sumit"
	list_interface := []interface{}{"sumit", 7}
	actual = validate.CountPresence("key", "sumit", list_interface)
	assert.Equal(true, actual)
	assert.Equal(false, validate.CountPresence("key", 100, list_interface))

}

func TestMaxLengthOfString(t *testing.T) {
	assert := assert.New(t)

	validate := Validate{}
	assert.NotNil(validate)

	maxLen := validate.MaxLengthOfString("key", 10, "valuevaluevaluevalue")
	assert.False(maxLen)
	expected := "key: Max allowed length is 10, found 20"
	actualErrList, _ := validate.ErrorList.Find("key")
	actual := actualErrList.Errors[0].Error()
	assert.Equal(expected, actual)
}

func TestMinLengthOfString(t *testing.T) {
	assert := assert.New(t)

	validate := Validate{}
	assert.NotNil(validate)

	minLen := validate.MinLengthOfString("key", 10, "value")
	assert.False(minLen)
	expected := "key: Min required length is 10, found 5"
	actualErrList, _ := validate.ErrorList.Find("key")
	actual := actualErrList.Errors[0].Error()
	assert.Equal(expected, actual)
}

func TestMatchRegExp(t *testing.T) {
	assert := assert.New(t)

	validate := Validate{}
	assert.NotNil(validate)

	formatCheck := validate.MatchRegExp("phone", "^([0-9])+$", "+91-9000000000")
	if assert.False(formatCheck) {
		expected := "phone: Format(^([0-9])+$) doesn't match with (+91-9000000000)"
		actualErrList, _ := validate.ErrorList.Find("phone")
		actual := actualErrList.Errors[0].Error()
		assert.Equal(expected, actual)
	}

	validate = Validate{}
	formatCheck = validate.MatchRegExp("phone", "^([0-9])+$", "+919000000000")
	if assert.False(formatCheck) {
		expected := "phone: Format(^([0-9])+$) doesn't match with (+919000000000)"
		actualErrList, _ := validate.ErrorList.Find("phone")
		actual := actualErrList.Errors[0].Error()
		assert.Equal(expected, actual)
	}

	validate = Validate{}
	formatCheck = validate.MatchRegExp("phone", "^([0-9])+$", "00919000000000")
	assert.True(formatCheck)
	actualErrList, _ := validate.ErrorList.Find("phone")
	actual := len(actualErrList.Errors)
	assert.Equal(0, actual)
}

func TestMatchEmail(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		Email    string
		Expected string
	}{
		{
			"sui@gm",
			"email: Email format is not correct (sui@gm)",
		},
		{
			"sui",
			"email: Email format is not correct (sui)",
		},
	}

	for _, e := range tests {
		validate := Validate{}
		assert.NotNil(validate)
		var formatCheck = validate.MatchEmail("email", e.Email)
		if assert.False(formatCheck) {
			actualErrList, _ := validate.ErrorList.Find("email")
			actual := actualErrList.Errors[0].Error()
			assert.Equal(e.Expected, actual)
		}
	}

	validate := Validate{}
	assert.NotNil(validate)

	var formatCheck = validate.MatchEmail("email", "sumitasok@gmail.com")
	assert.True(formatCheck)

}
