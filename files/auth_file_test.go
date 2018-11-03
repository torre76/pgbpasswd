package files

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFileExists(t *testing.T) {
	assert := assert.New(t)
	fileManager := AuthFileFileManager{}

	assert.True(fileManager.fileExists("/etc/sudoers"))
	assert.False(fileManager.fileExists("/etc/bashrca"))
}

func TestCopyFile(t *testing.T) {
	assert := assert.New(t)
	fileManager := AuthFileFileManager{}

	assert.Nil(fileManager.copyFile("/etc/localtime", "/tmp/localtime"))

}
