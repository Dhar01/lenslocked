package main

import (
	"html/template"
	"os"
)

type User struct {
	Name    string
	Age     int
	Address string
}

func main() {

	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}
	user := User{
		Name:    "John Smith",
		Age:     22,
		Address: "Bangladesh",
	}
	err = t.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}

	// scenario
	// err := CreateUser()
	// if err != nil {
	// 	log.Println(err)
	// }
	// err = CreateOrg()
	// if err != nil {
	// 	log.Println(err)
	// }

}

// // replace
// func Connect() error {
// 	return errors.New("connection failed")
// }
// func CreateUser() error {
// 	err := Connect()
// 	if err != nil {
// 		return fmt.Errorf("create user: %w", err)
// 	}
// 	return nil
// }
// func CreateOrg() error {
// 	err := CreateUser()
// 	if err != nil {
// 		return fmt.Errorf("create org: %w", err)
// 	}
// 	return nil
// }
