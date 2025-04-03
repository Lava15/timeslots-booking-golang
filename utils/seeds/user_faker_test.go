package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateFakeUsers(t *testing.T) {
	n := 10
	users := GenerateFakeUsers(n)
	assert.Len(t, users, n, "should generate the requested number of users")
	for _, user := range users {
		assert.NotEmpty(t, user.Name, "user name should not be empty")
		assert.NotEmpty(t, user.Email, "user email should not be empty")
		assert.NotEmpty(t, user.Password, "user password should not be empty")
	}
}
