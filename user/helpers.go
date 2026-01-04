package user

func getUserById(users []User, id string) *User {
	for i := range users {
		if users[i].Id == id {
			return &users[i]
		}
	}
	return nil
}
