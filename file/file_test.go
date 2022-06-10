package file

import (
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestFindByExtension(t *testing.T) {

	afs := &afero.Afero{Fs: afero.NewMemMapFs()}

	afs.WriteFile("example.txt", []byte("test"), 0644)
	afs.WriteFile("example.csv", []byte("test"), 0644)
	afs.WriteFile("example2.csv", []byte("test"), 0644)
	result, err := FindByExtension(afs, ".", "txt", "csv")
	assert.Equal(t, 3, len(result))
	assert.NoError(t, err)

	afs.Mkdir("temp", 0644)
	afs.WriteFile("temp/example.txt", []byte("test"), 0644)
	result, err = FindByExtension(afs, "temp", "txt", "csv")
	assert.Equal(t, 1, len(result))
	assert.NoError(t, err)
}

func Test_generateRegex(t *testing.T) {
	result, err := generateRegex("txt", "csv")
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, ".*\\.(txt|csv)", result.String())

	_, err = generateRegex("^\\/(?!\\/)(.*?)", "csv")
	assert.Error(t, err)
}
