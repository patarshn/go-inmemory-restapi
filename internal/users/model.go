package users

type UserModel struct {
	ID       *int64 `form:"id"`
	Username string `form:"username"`
}
