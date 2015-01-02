package validator

import (
	// "fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFieldRequired(t *testing.T) {
	assert := assert.New(t)

	tr := TestRequired{"sumit", "password"}

	result := Parse(tr)

	assert.Empty(result)
}

type TestRequired struct {
	Username string `validator:"required,anotherValidator"`
	Password string `validator:"required"`
}
