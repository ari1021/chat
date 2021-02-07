package request

type CreateRoom struct {
	Name string `form:"name" validate:"required,excludesall= "`
}

type DeleteRoom struct {
	ID uint `param:"id"`
}
