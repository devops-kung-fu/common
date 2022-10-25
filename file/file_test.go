package file

import (
	"path/filepath"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

const BAD_REGEX = "^\\/(?!\\/)(.*?)"

func TestSHA(t *testing.T) {
	afs := &afero.Afero{Fs: afero.NewMemMapFs()}
	f, _ := filepath.Abs("example.txt")
	afs.WriteFile(f, []byte("test"), 0644)

	sha, err := SHA(afs, "example.txt")
	assert.Equal(t, "9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08", sha)
	assert.NotNil(t, sha)
	assert.NoError(t, err)

	_, err = SHA(afs, "nofile")
	assert.Error(t, err)
}

func TestFindByRegex(t *testing.T) {
	afs := &afero.Afero{Fs: afero.NewMemMapFs()}

	afs.WriteFile("example.txt", []byte("test"), 0644)
	files, err := FindByRegex(afs, ".", "(.*)\\.txt")
	assert.Len(t, files, 1)
	assert.NoError(t, err)

	_, err = FindByRegex(afs, ".", BAD_REGEX)
	assert.Error(t, err)
}

func TestFindByExtension(t *testing.T) {

	afs := &afero.Afero{Fs: afero.NewMemMapFs()}

	afs.WriteFile("example.txt", []byte("test"), 0644)
	afs.WriteFile("example.csv", []byte("test"), 0644)
	afs.WriteFile("example2.csv", []byte("test"), 0644)
	result, err := FindByExtension(afs, ".", []string{"txt", "csv"})
	assert.Equal(t, 3, len(result))
	assert.NoError(t, err)

	result, err = FindByExtension(afs, ".", []string{BAD_REGEX, "csv"})
	assert.Error(t, err)

	afs.Mkdir("temp", 0644)
	afs.WriteFile("temp/example.txt", []byte("test"), 0644)
	result, err = FindByExtension(afs, "temp", []string{"txt", "csv"})
	assert.Equal(t, 1, len(result))
	assert.NoError(t, err)
}

func Test_generateRegex(t *testing.T) {
	result, err := generateRegex([]string{"txt", "csv"})
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, ".*\\.(txt|csv)", result.String())

	_, err = generateRegex([]string{BAD_REGEX, "csv"})
	assert.Error(t, err)
}
