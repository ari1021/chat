package model

type User struct {
	ID   string
	Name string `form:"name" validate:"required"`
}

var Users = map[string]*User{}
