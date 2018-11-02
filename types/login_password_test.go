package types

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoginPasswordStruct(t *testing.T) {
	assert := assert.New(t)

	lp := NewLoginPassword("pippo", "pippo")
	lpHashed := NewLoginHashedPassword("pippo", "md5e8db5c992bd46882190967eb213a233c")

	assert.Equal(lp, lpHashed)

	fmt.Printf("LoginPassword: %s\n", lp)
}
