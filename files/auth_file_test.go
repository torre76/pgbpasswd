package files

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFileExists(t *testing.T) {

	assert := assert.New(t)
	fileManager := AuthFileFileManager{}

	assert.True(fileManager.fileExists("/etc/bashrc"))
	assert.False(fileManager.fileExists("/etc/bashrca"))
}
