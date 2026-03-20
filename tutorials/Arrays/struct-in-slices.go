package arrays

import (
	"strings"
	"test/model"
)

func FindUserNames(users []model.User) []string {

	names := make([]string, 0)

	for _, user := range users {

		_, emailBox, _ := strings.Cut(user.Email, "@")

		if emailBox == "example.com" {
			names = append(names, user.Name)
		}
	}

	if len(names) == 0 {
		return names
	}

	return names
}
