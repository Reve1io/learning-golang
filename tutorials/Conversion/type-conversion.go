package conversion

import (
	"fmt"
	"strconv"
)

func BuildProfile(name string, age int, rating float64) string {
	strAge := strconv.Itoa(age)
	strRating := fmt.Sprintf("%.1f", rating)

	result := "Name: " + name + ", Age: " + strAge + ", Rating: " + strRating

	return result
}
