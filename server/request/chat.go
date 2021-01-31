package request

type GetChats struct {
	RoomID int `param:"id"`
	Limit  int `query:"limit" validate:"required"`
	Offset int `query:"offset"`
}

type CreateChat struct {
	Message  string `form:"message" validate:"required,excludesall= "`
	RoomID   int    `param:"id"`
	UserName string `form:"user_name" validate:"required,excludesall= "`
}
