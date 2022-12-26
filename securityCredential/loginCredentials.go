package securityCredential

type LoginCredentials struct {
	Email    string `form:"email"`
	Phone    string `from:"phone"`
	Password string `form:"password"`
}
