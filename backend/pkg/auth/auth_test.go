package auth

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestAuth(t *testing.T) {
	t.Run("genToken", func(t *testing.T) {
		jwt := NewJWT("secret")
		userId := 32423423
		token, err := jwt.GenerateToken(uint64(userId))

		assert.Equal(t, err, nil)
		assert.NotEqual(t, token, "")
		t.Logf("token: %v", token)

		c, err := jwt.ParseToken(token)
		assert.Equal(t, err, nil)
		assert.NotEqual(t, c, nil)
		t.Logf("userId: %v c.UserId: %v", userId, c.UserID)
		assert.IsEqual(c.UserID, uint64(userId))

	})

	t.Run("parseInvalidToken", func(t *testing.T) {
		jwt := NewJWT("secret")
		token := "invalidToken"
		_, err := jwt.ParseToken(token)
		t.Errorf("parseInvalidToken err: %v", err)
		assert.NotEqual(t, err, nil)
	})
}
