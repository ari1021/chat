package model

import "gopkg.in/go-playground/validator.v9"

type User struct {
	ID   string
	Name string `form:"name" validate:"required"`
}

type UserValidator struct {
	validator *validator.Validate
}

func (uv *UserValidator) Validate(user *User) error {
	return uv.validator.Struct(user)
}

var Users = map[string]*User{}
