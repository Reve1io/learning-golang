package model

type User struct {
	Name  string
	Email string
}

type Order struct {
	CustomerID int
	Price      int
}
