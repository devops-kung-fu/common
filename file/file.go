// Package file contains functions for finding files.
package file

import (
	"os"
	"regexp"
	"strings"

	"github.com/spf13/afero"
)

// FindByRegex recursively searches for files matching a pattern.
func FindByRegex(afs *afero.Afero, root string, re string) (files []string, err error) {
	libRegEx, err := regexp.Compile(re)
	if err != nil {
		return
	}
	err = afero.Walk(afs, root, func(filePath string, info os.FileInfo, err error) error {
		if err == nil && libRegEx.MatchString(info.Name()) {
			files = append(files, filePath)
			return nil
		}
		return err
	})

	return
}

// FindByExtension walks a file tree and returns a list of files based on extensions passed
func FindByExtension(afs *afero.Afero, path string, extensions []string) (files []string, err error) {

	regex, err := generateRegex(extensions)
	if err != nil {
		return
	}

	err = afs.Walk(path, func(filePath string, info os.FileInfo, err error) error {
		if err == nil && regex.MatchString(info.Name()) {
			files = append(files, filePath)
			return nil
		}
		return err
	})

	return
}

// func generateRegex(extensions []string) (regex *regexp.Regexp, err error) {
// 	re := ".*\\.("
// 	for _, val := range extensions {
// 		re += fmt.Sprintf("%s|", val)
// 	}
// 	re = strings.TrimSuffix(re, "|")
// 	re += string(")")

// 	regex, err = regexp.Compile(re)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return regex, nil
// }

func generateRegex(extensions []string) (*regexp.Regexp, error) {
	re := ".*\\.(" + strings.Join(extensions, "|") + ")"
	regex, err := regexp.Compile(re)
	if err != nil {
		return nil, err
	}
	return regex, nil
}
