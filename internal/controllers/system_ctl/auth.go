package system_ctl

import (
	"github.com/gin-gonic/gin"
	"github.com/zs368/gin-example/internal/models"
	"github.com/zs368/gin-example/internal/pkg/app"
	"github.com/zs368/gin-example/internal/pkg/errcode"
	"github.com/zs368/gin-example/internal/rules/system_rule"
	"github.com/zs368/gin-example/pkg/database"
)

type Auth struct{}

func NewAuth() Auth {
	return Auth{}
}

func (a Auth) Login(c *gin.Context) {
	var (
		r      = app.NewResponse(c)
		params = system_rule.LoginRequest{}
	)

	if err := app.BindAndValid(c, &params); err != nil {
		r.ToErrorResponse(errcode.InvalidParams.WithDetails(err.Errors()...))
		return
	}

	var (
		db = database.DB
		u  = struct {
			Id   uint   `json:"id"`
			Name string `json:"name"`
			Role string `json:"role"`
		}{}
		user models.User
	)

	if err := db.Model(&user).Where("email = ? AND password = ?", params.Username, params.Password).First(&u).Error; err != nil {
		r.ToErrorResponse(errcode.UnauthorizedAuthNotExist)
		return
	}

	token, err := app.GenerateToken(app.UserInfo{
		Uid:      1,
		Username: "zs",
		Role:     "admin",
	})

	if err != nil {
		r.ToErrorResponse(errcode.UnauthorizedTokenGenerate)
		return
	}

	r.ToResponse(system_rule.LoginResponse{Token: token})
}
