package encrypt

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

// PgMd5HashedPassword builds the *md5 like* representation for a PostgreSQL password.
func PgMd5HashedPassword(login string, password string) string {
	hasher := md5.New()
	hasher.Write([]byte(strings.TrimSpace(password) + strings.TrimSpace(login)))
	return "md5" + hex.EncodeToString(hasher.Sum(nil))
}
