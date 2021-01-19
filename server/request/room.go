package request

type CreateRoom struct {
	Name   string `form:"name" validate:"required,excludesall= "`
	UserId int    `form:"user_id" validate:"required"`
}
