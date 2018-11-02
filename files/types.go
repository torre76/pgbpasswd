package files

import "github.com/torre76/pgbpasswd/types"

// FileManager writes a collection of Login and Hashed Password to a file
// Since there are at least two different formats there will be multiple implementations
type FileManager interface {
	fileExists(filename string) (bool, error)

	removeFile(fileName string) error

	Read(fileName string) ([]types.LoginPassword, error)

	Write(fileName string, elements []types.LoginPassword) error
}

// baseFileManager is a base implementation for common method used by FileManager
type baseFileManager struct {
	FileManager
}

func (fm *baseFileManager) fileExists(filename string) (bool, error) {
	return false, nil
}

func (fm *baseFileManager) removeFile(filename string) error {
	return nil
}
