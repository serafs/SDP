package model

type User struct {
	Username string
}

type UserService interface {
	Create(user *User) (*User, error)
	Get(userId string) (*User, error)
}

type UserRepository interface {
	Create(user *User) (*User, error)
	Get(userId string) (*User, error)
}
