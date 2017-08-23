package user

func getUser(where map[string]interface{}) (*User, error) {
	var u User
	return &u, TableUser.ReadOne(where, &u)
}

func GetByID(id string) (*User, error) {
	var u User
	return &u, TableUser.ReadByID(id, &u)
}

func GetByUsername(username string) (*User, error) {
	var u User
	return &u, TableUser.ReadOne(map[string]interface{}{
		"username": username,
		"dtime":    0,
	}, &u)
}
func GetByEmail(email string) (*User, error) {
	var u User
	return &u, TableUser.ReadOne(map[string]interface{}{
		"email": email,
		"dtime": 0,
	}, &u)
}

func GetAll() ([]*User, error) {
	var users = []*User{}
	return users, TableUser.UnsafeReadAll(&users)
}
