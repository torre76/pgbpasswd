package files

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFileExists(t *testing.T) {
	assert := assert.New(t)
	fileManager := NewAuthFileFileManager()

	assert.True(fileManager.fileExists("/etc/sudoers"))
	assert.False(fileManager.fileExists("/etc/bashrca"))
}

func TestCopyFile(t *testing.T) {
	assert := assert.New(t)
	fileManager := NewAuthFileFileManager()

	assert.Nil(fileManager.copyFile("/etc/localtime", "/tmp/localtime"))
	assert.True(fileManager.fileExists("/tmp/localtime"))
}

func TestRemoveFile(t *testing.T) {
	assert := assert.New(t)
	fileManager := NewAuthFileFileManager()

	assert.Nil(fileManager.copyFile("/etc/localtime", "/tmp/localtime_to_delete"))
	assert.True(fileManager.fileExists("/tmp/localtime_to_delete"))

	assert.Nil(fileManager.removeFile("/tmp/localtime_to_delete"))
	assert.False(fileManager.fileExists("/tmp/localtime_to_delete"))

}
