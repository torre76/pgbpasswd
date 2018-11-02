package files

import "github.com/torre76/pgbpasswd/types"

// AuthFileFileManager is a FileManager that writes PGBouncer Auth File Style files
type AuthFileFileManager struct{}

func (fm *AuthFileFileManager) Read(fileName string) ([]types.LoginPassword, error) {
	return nil, nil
}

func (fm *AuthFileFileManager) Write(fileName string, elements []types.LoginPassword) error {
	return nil
}
