package model

type User struct {
	ID   string
	Name string `form:"name"`
}

var Users = map[string]*User{}
