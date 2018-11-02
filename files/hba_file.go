package files

import "github.com/torre76/pgbpasswd/types"

// HBAFileFileManager is a FileManager that writes Postgres HBA Auth File Style files
type HBAFileFileManager struct {
	baseFileManager
}

func (fm *HBAFileFileManager) Read(fileName string) ([]types.LoginPassword, error) {
	return nil, nil
}

func (fm *HBAFileFileManager) Write(fileName string, elements []types.LoginPassword) error {
	return nil
}
