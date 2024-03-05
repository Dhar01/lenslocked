package main

import (
	"fmt"
	// "html/template"
	// "os"
)

type User struct {
	Name    string
	Age     int
	Address string
}

func main() {

	// t, err := template.ParseFiles("hello.gohtml")
	// if err != nil {
	// 	panic(err)
	// }
	// user := User{
	// 	Name:    "John Smith",
	// 	Age:     22,
	// 	Address: "Bangladesh",
	// }
	// err = t.Execute(os.Stdout, user)
	// if err != nil {
	// 	panic(err)
	// }

	Demo()
	Demo(1)
	Demo(1, 2, 3)


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

func Demo(numbers ...int) {
	for _, number := range numbers {
		fmt.Print(number, " ")
	}
	fmt.Println()
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
