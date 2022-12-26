package service

type LoginService interface {
	LoginUser(email string, phone string, password string) bool
}
type loginInformation struct {
	email    string
	phone    string
	password string
}

func StaticLoginService() LoginService {
	return &loginInformation{
		email:    "info@gmail.com",
		phone:    "+964422",
		password: "123456789",
	}
}
func (info *loginInformation) LoginUser(email string, phone string, password string) bool {
	return info.email == email && info.phone == phone && info.password == password
}
