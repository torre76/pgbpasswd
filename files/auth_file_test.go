package files

import (
	"path/filepath"
	"testing"

	"github.com/torre76/pgbpasswd/types"

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

func TestReadMissingFile(t *testing.T) {
	assert := assert.New(t)
	fileManager := NewAuthFileFileManager()

	_, err := fileManager.Read("asset")

	assert.Equal(err.Error(), "Cannot read from file, does it exists?")
}

func TestReadMalformed(t *testing.T) {
	assert := assert.New(t)
	fileManager := NewAuthFileFileManager()
	asset, _ := filepath.Abs("../test_data/users_malformed.txt")

	_, err := fileManager.Read(asset)

	assert.Equal(err.Error(), "File format is not valid")
}

func TestWrite(t *testing.T) {
	assert := assert.New(t)
	fileManager := NewAuthFileFileManager()
	tmp, _ := filepath.Abs("../test_data/users_write.txt")
	fileManager.removeFile(tmp)

	var mockData = []types.LoginPassword{
		*types.NewLoginPassword("pippo", "pippo"),
		*types.NewLoginPassword("astro", "astro"),
		*types.NewLoginPassword("dae\"mon", "astro"),
	}

	err := fileManager.Write(tmp, mockData)
	assert.Nil(err)
}
