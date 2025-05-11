package model

type User struct {
	ID       string
	Username string
	Email    string
	Phone    string
	Password string
}

type UserWithMeta struct {
	User
	IpAddress string
}
