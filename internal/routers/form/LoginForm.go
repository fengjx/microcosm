package form

type LoginForm struct {
	Username string `form:"username"`
	Password string `form:"password"`
	Captcha  string `form:"captcha"`
}
