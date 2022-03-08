package controller

import (
	"net/http"
	"task/api/service"
	"task/models"
	"task/util"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type UserController struct {
	service service.UserService
}

func NewUserController(s service.UserService) UserController {
	return UserController{
		service: s,
	}
}

func (u *UserController) CreateUser(c *gin.Context) {
	var user models.UserRegister
	if err := c.ShouldBind(&user); err != nil {
		util.ErrorJSON(c, http.StatusBadRequest, "Inavlid Json Provided")
		return
	}

	hashPassword, _ := util.HashPassword(user.Password)
	user.Password = hashPassword

	err := u.service.CreateUser(user)
	if err != nil {
		util.ErrorJSON(c, http.StatusBadRequest, "Failed to create user")
		return
	}

	util.SuccessJSON(c, http.StatusOK, "Successfully Created user")
}

func (u *UserController) Login(c *gin.Context) {
	var user models.UserLogin
	var hmacSampleSecret []byte
	if err := c.ShouldBindJSON(&user); err != nil {
		util.ErrorJSON(c, http.StatusBadRequest, "Inavlid Json Provided")
		return
	}
	dbUser, err := u.service.LoginUser(user)
	if err != nil {
		util.ErrorJSON(c, http.StatusBadRequest, "Invalid Login Credentials")
		return
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": dbUser,
		"exp":  time.Now().Add(time.Minute * 15).Unix(),
	})

	tokenString, err := token.SignedString(hmacSampleSecret)
	if err != nil {
		util.ErrorJSON(c, http.StatusBadRequest, "Failed to get token")
		return
	}
	response := &util.Response{
		Success: true,
		Message: "Token generated sucessfully",
		Data:    tokenString,
	}
	c.JSON(http.StatusOK, response)
}
