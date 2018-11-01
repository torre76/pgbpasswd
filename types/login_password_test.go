package types

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoginPasswordStruct(t *testing.T) {
	assert := assert.New(t)

	lp := NewLoginPassword("pippo", "pippo")

	assert.Equal(lp.Login, "pippo")
	assert.Equal(lp.HashedPassword, "md5e8db5c992bd46882190967eb213a233c")

	fmt.Printf("LoginPassword: %s\n", lp)
}
