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
	result, err := FindByExtension(afs, ".", []string{"txt", "csv"})
	assert.Equal(t, 3, len(result))
	assert.NoError(t, err)

	result, err = FindByExtension(afs, ".", []string{"^\\/(?!\\/)(.*?)", "csv"})
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

	_, err = generateRegex([]string{"^\\/(?!\\/)(.*?)", "csv"})
	assert.Error(t, err)
}
