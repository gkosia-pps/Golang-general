package main

import "fmt"

type User struct {
	name string
}

func getUser(name string) (User, error) {

	if name != "" {

		return User{
			name: "Gab",
		}, nil
	} else {
		return User{}, fmt.Errorf("No username passed")
	}
}

func main() {

	// handle error
	u, err := getUser("")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(u.name)

}
