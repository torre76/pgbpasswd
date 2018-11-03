package files

import "github.com/torre76/pgbpasswd/types"

// HBAFileFileManager is a FileManager that writes Postgres HBA Auth File Style files
type hbaFileFileManager struct {
	baseFileManager
}

func (fm *hbaFileFileManager) Read(fileName string) ([]types.LoginPassword, error) {
	return nil, nil
}

func (fm *hbaFileFileManager) Write(fileName string, elements []types.LoginPassword) error {
	return nil
}

// NewHBAFileManager create an instance of File Manager that deals with with Postgres HBA Auth File format
func NewHBAFileManager() FileManager {
	return &hbaFileFileManager{}
}
