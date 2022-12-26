package api

import (
	"Gin-JWT-Sample/securityCredential"
	"Gin-JWT-Sample/service"
	"github.com/gin-gonic/gin"
)

type LoginAPI interface {
	Login(ctx *gin.Context) string
}

type loginAPI struct {
	loginService service.LoginService
	jWtService   service.JWTService
}

func LoginHandler(loginService service.LoginService,
	jWtService service.JWTService) LoginAPI {
	return &LoginAPI{
		loginService: loginService,
		jWtService:   jWtService,
	}
}

func (api *loginAPI) Login(ctx *gin.Context) string {
	var credential securityCredential.LoginCredentials
	err := ctx.ShouldBind(&credential)
	if err != nil {
		return "no data found"
	}
	isUserAuthenticated := service.LoginService.LoginUser(credential.Email, credential.Password)
	if isUserAuthenticated {
		return api.jWtService.GenerateToken(credential.Email, true)

	}
	return ""
}
