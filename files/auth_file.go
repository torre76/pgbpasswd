package files

import (
	"github.com/torre76/pgbpasswd/types"
)

// authFileFileManager is a FileManager that writes PGBouncer Auth File Style files
type authFileFileManager struct {
	baseFileManager
}

func (fm *authFileFileManager) Read(fileName string) ([]types.LoginPassword, error) {
	return nil, nil
}

func (fm *authFileFileManager) Write(fileName string, elements []types.LoginPassword) error {
	return nil
}

// NewAuthFileFileManager create an instance of File Manager that deals with with PgBouncer Auth File format
func NewAuthFileFileManager() FileManager {
	return &authFileFileManager{}
}
