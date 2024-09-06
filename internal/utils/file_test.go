package utils

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFileExists(t *testing.T) {
	tmp := t.TempDir()

	assert.True(t, FileExist(tmp))
	inexistent := filepath.Join(tmp, "inexistent")
	assert.False(t, FileExist(inexistent))
}

func TestMkRelDir(t *testing.T) {
	tmp := t.TempDir()
	os.Chdir(tmp)

	err := MkRelDir("dir")
	assert.NoError(t, err)
	assert.True(t, FileExist(filepath.Join(tmp, "dir")))
}
