package main

import (
	"fmt"
	"test/model"
	arrays "test/tutorials/Arrays"
)

func main() {
	users := []model.User{
		{Name: "Семен", Email: "semen@example.com"},
		{Name: "Юля", Email: "yulia@gmail.com"},
		{Name: "Алексей", Email: "alex@example.com"},
	}

	result := arrays.FindUserNames(users)
	fmt.Println(result)
}
