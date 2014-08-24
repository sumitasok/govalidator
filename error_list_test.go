package validator

import (
	// "fmt"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	item_err_msg = "Max allowed charaters are 50"
)

func GetErrorList() ErrorList {
	return ErrorList{
		Errors: []ParamsError{
			ParamsError{
				Field: "item",
				Errors: []error{
					errors.New(item_err_msg),
				},
			},
		},
	}
}

func TestFindIndex(t *testing.T) {
	assert := assert.New(t)

	err_list := GetErrorList()

	erridx, err := err_list.FindIndex("item")
	assert.Nil(err)
	errmsg := err_list.Errors[erridx].Errors
	assert.Equal(item_err_msg, errmsg[0].Error())
	assert.Equal(0, erridx)

	_, err1 := err_list.FindIndex("items")
	assert.NotNil(err1)

}

func TestAppend(t *testing.T) {
	assert := assert.New(t)

	err_list := GetErrorList()

	err_list.Append("item", errors.New("a new error"))
	txn_err_idx, txnfinderr := err_list.FindIndex("item")
	if assert.Nil(txnfinderr) {
		assert.Equal(2, len(err_list.Errors[txn_err_idx].Errors))
		assert.Equal(item_err_msg, err_list.Errors[txn_err_idx].Errors[0].Error())
		assert.Equal("a new error", err_list.Errors[txn_err_idx].Errors[1].Error())
	}

	// Adding a new key value pair

	err_list.Append("newid", errors.New("a new error"))
	txn_err_idx1, txnfinderr1 := err_list.FindIndex("newid")
	if assert.Nil(txnfinderr1) {
		assert.Equal(1, len(err_list.Errors[txn_err_idx1].Errors))
		assert.Equal("a new error", err_list.Errors[txn_err_idx1].Errors[0].Error())
	}

}
