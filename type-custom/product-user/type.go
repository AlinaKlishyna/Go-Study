package main

import "fmt"

type UserID int

type User struct {
	ID        UserID
	FirstName string
}

func (u UserID) String() string {
	return fmt.Sprintf("UserID-%d", u)
}

//-----------------------------

type ProductID int

type Product struct {
	ID    ProductID
	Title string
}

func (p ProductID) String() string {
	return fmt.Sprintf("ProductID-%d", p)
}

// -----------------------------
func main() {
	user := User{
		ID:        1,
		FirstName: "Alyona",
	}

	product := Product{
		ID:    10,
		Title: "Apple",
	}

	fmt.Println(user)
	fmt.Println(product)

}
