package main

import (
	"fmt"
)

// Error is an interface in go with a method signuture Error() string
// I can implement the interface to create custom errors
type customUserError struct {
	name string
}

func (ue customUserError) Error() string {
	return fmt.Sprintf("Error for user %v", ue.name)
}

type User struct {
	name string
}

func getUserWithError() error {
	return customUserError{
		name: "AAAAAA",
	}
}
func getUser(name string) (User, error) {

	if name != "" {

		return User{
			name: "Gab",
		}, nil

	} else {
		return User{}, fmt.Errorf("No username passed")
		// OR use errorsw package
		//return User{}, errors.New("This is an error from package")
	}
}

func main_08() {

	// handle error
	u, err := getUser("")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(u.name)

	// use the custom error
	// custom errors allows us to store extra information that we can use to format the error message
	err = getUserWithError()
	if err != nil {
		fmt.Println(err)
	}

}
