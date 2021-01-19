package request

type CreateRoom struct {
	Name string `form:"name" validate:"required,excludesall= "`
}
