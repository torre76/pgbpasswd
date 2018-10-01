package encrypt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPgMd5HashedPassword(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(PgMd5HashedPassword("pippo", "pippo"), "md5e8db5c992bd46882190967eb213a233c", "Wrong MD5 PgBouncer password built")
	assert.Equal(PgMd5HashedPassword("gestionale_lead", "DhZ29NPf"), "md59b20b9d542d4a2e0866b5cf6c9234468", "Wrong MD5 PgBouncer password built")
}
