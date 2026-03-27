package maps

import (
	"errors"
	"strconv"
)

func GetGrade(scores map[string]int, name string) (string, error) {

	if scores == nil {
		return "", errors.New("Map is empty")
	}

	if name == "" {
		return "", errors.New("Name is empty")
	}

	score, err := scores[name]
	if err == false {
		return "", errors.New("Non-existing student")
	}

	str := name + " has " + strconv.Itoa(score) + " points"

	return str, nil
}
