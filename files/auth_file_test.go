package files

import (
	"path/filepath"
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

func TestRead(t *testing.T) {
	assert := assert.New(t)
	fileManager := NewAuthFileFileManager()
	asset, _ := filepath.Abs("../test_data/users.txt")

	result, _ := fileManager.Read(asset)
	assert.NotNil(result)
	assert.True(len(result) == 3)
	assert.Equal(result[2].Login, `macha_"test"`)
}
