package files

import (
	"errors"
	"io"
	"os"

	"github.com/torre76/pgbpasswd/types"
)

// FileManager writes a collection of Login and Hashed Password to a file
// Since there are at least two different formats there will be multiple implementations
type FileManager interface {
	// fileExists checks if a file exists on filesystem
	fileExists(filename string) bool

	removeFile(fileName string) error

	// copyFile copy a file from source to destination
	copyFile(srcFileName string, destFileName string) error

	Read(fileName string) ([]types.LoginPassword, error)

	// Write create a file into fylesystem that contains login and hashed password using a specified file format
	Write(fileName string, elements []types.LoginPassword) error
}

// baseFileManager is a base implementation for common method used by FileManager
type baseFileManager struct {
	FileManager
}

func (fm *baseFileManager) copyFile(srcFileName string, destFileName string) error {
	if !fm.fileExists(srcFileName) {
		return errors.New("Source File does not exists")
	}

	source, err := os.Open(srcFileName)
	if err != nil {
		return err
	}
	defer source.Close()

	destination, err := os.Create(destFileName)
	if err != nil {
		return err
	}
	defer destination.Close()

	bytes, err := io.Copy(destination, source)
	_ = bytes
	if err != nil {
		return err
	}

	return nil
}

func (fm *baseFileManager) fileExists(filename string) bool {
	if _, err := os.Stat(filename); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func (fm *baseFileManager) removeFile(filename string) error {
	return nil
}
