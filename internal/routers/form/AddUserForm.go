package form

type AddUserForm struct {
	Username string `form:"username" binding:"required"`
	Email    string `form:"email"`
	Mobile   string `form:"mobile"`
}
