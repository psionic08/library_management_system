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

func RemoveUser(id string) error {
	for i, u := range users {
		if u.Id == id {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("book with id: %s does not exist", id)
}

func GetUserById(id string) (*User, error) {
	existingUser := getUserById(users, id)
	if existingUser == nil {
		return nil, fmt.Errorf("user with id: %s does not exist", id)
	}
	clone := *existingUser
	return &clone, nil
}
