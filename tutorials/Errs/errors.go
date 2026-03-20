package errs

import (
	"errors"
	"strings"
)

func GetFileExtension(filename string) (string, error) {

	var sep rune

	for _, v := range filename {

		if v == '.' {
			sep = v
		}
	}

	if sep == 0 {
		return "", errors.New("separate undefined!")
	}

	extensions := strings.Split(filename, ".")

	extension := extensions[len(extensions)-1]

	return extension, nil
}
