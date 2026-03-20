package strings

import "strconv"

func BuildGreeting(name string, count int) string {
	msg := "Hello, " + name + "! You have " + strconv.Itoa(count) + " new messages."
	return msg
}

func GetHiddenCard(number string, count int) string {
	lastNum := number[12:]
	stars := "************"
	lastWithStars := stars[:count]

	result := lastWithStars + lastNum

	return result
}
