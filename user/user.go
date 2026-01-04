package user

import "fmt"

type User struct {
	Id   string
	Name string
}

var users []User

func AddUser(id string, name string) error {
	existingUser := getUserById(users, id)
	if existingUser != nil {
		return fmt.Errorf("user with id %s already exists", id)
	}
	u := &User{
		id,
		name,
	}

	users = append(users, *u)

	return nil
}

func GetUsers() []User {
	output := make([]User, len(users))
	copy(output, users)
	return output
}
