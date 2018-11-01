package types

import (
	"fmt"

	"github.com/torre76/pgbpasswd/encrypt"
)

// LoginPassword contains Postgres login and hashed password that is managed by the application
type LoginPassword struct {
	Login          string
	HashedPassword string
}

func (lp *LoginPassword) String() string {
	return fmt.Sprintf("Login: [%q] - Hashed Password: [%q]", lp.Login, lp.HashedPassword)
}

// NewLoginPassword initialize a LoginPassword structure starting from Postgres login and password
func NewLoginPassword(login string, password string) *LoginPassword {
	return &LoginPassword{login, encrypt.PgMd5HashedPassword(login, password)}
}
