package data

type User struct {
	name string
}

func NewUserDao(usn string) User {
	return User{name: usn}
}

func (u *User) GetUserName() string {
	return u.name
}
