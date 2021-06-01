package bcrypt

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBcrypt(t *testing.T) {
	// Given
	password := "test_password"

	// When
	hashedPass, err := HashPassword(password)

	// Then
	assert.NoError(t, err)
	assert.True(t, IsPasswordHash(password, hashedPass))
}
