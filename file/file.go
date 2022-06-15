package file

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/spf13/afero"
)

func FindByExtension(afs *afero.Afero, path string, extensions ...any) (files []string, err error) {

	regex, err := generateRegex(extensions)
	if err != nil {
		return nil, err
	}

	err = afs.Walk(path, func(filePath string, info os.FileInfo, err error) error {
		if err == nil && regex.MatchString(info.Name()) {
			files = append(files, filePath)
		} else {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	return files, err
}

func generateRegex(extensions ...any) (regex *regexp.Regexp, err error) {

	re := ".*\\.("
	for _, val := range extensions {
		re += fmt.Sprintf("%s|", val)
	}
	re = strings.TrimSuffix(re, "|")
	re += string(")")

	regex, err = regexp.Compile(re)
	if err != nil {
		return nil, err
	}
	return regex, nil
}
