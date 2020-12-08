package model

type User struct {
	ID   string
	Name string `json."name"`
}

var Users map[string]*user
