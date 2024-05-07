package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Doni0414/go-canteen-project/models"
	"github.com/Doni0414/go-canteen-project/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var jwtKey = []byte("my_secret_key")

func AuthPage(c *gin.Context) {
	user, exists := c.Get("user")
	fmt.Println(user)
	if !exists {
		c.HTML(http.StatusOK, "login.html", gin.H{"title": "Логин/Регистрация"})
		return
	}

	c.HTML(http.StatusOK, "login.html", gin.H{"title": "Логин/Регистрация", "user": user})
}

func Auth(c *gin.Context) {
	authType := c.PostForm("auth-type")
	fmt.Println(authType)
	if authType == "login" {
		Login(c)
	} else if authType == "signup" {
		Signup(c)
	}
}
func Login(c *gin.Context) {
	var loginUser models.LoginUser

	if err := c.Bind(&loginUser); err != nil {
		c.HTML(400, "login.html", gin.H{"error": "400: Bad Request", "loginUser": loginUser})
		return
	}

	var existingUser models.User

	models.DB.Where("email = ?", loginUser.Email).First(&existingUser)

	if existingUser.ID == 0 {
		c.HTML(400, "login.html", gin.H{"error": "Пользователь не найден!", "loginUser": loginUser})
		return
	}

	if ok := utils.CompareHashPassword(loginUser.Password, existingUser.Password); !ok {
		c.HTML(400, "login.html", gin.H{"error": "Неправильный пароль!", "loginUser": loginUser})
		return
	}

	expirationTime := time.Now().Add(5 * time.Minute)

	claims := models.Claims{
		Role: existingUser.Role,
		StandardClaims: jwt.StandardClaims{
			Subject:   existingUser.Email,
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		c.HTML(500, "login.html", gin.H{"error": "Ошибка сервера!"})
		return
	}

	c.SetCookie("token", tokenString, int(expirationTime.Unix()), "/", "localhost", false, true)

	if existingUser.Role == "user" {
		c.Redirect(302, "/user/main")
	} else {
		c.Redirect(302, "/employee/main")
	}
}

func Signup(c *gin.Context) {
	var signupUser models.SignupUser

	if err := c.Bind(&signupUser); err != nil {
		c.HTML(400, "login.html", gin.H{"error": "400: Bad Request", "signupUser": signupUser})
		return
	}

	var existingUser models.User

	models.DB.Where("email = ?", signupUser.Email).First(&existingUser)

	if existingUser.ID != 0 {
		c.HTML(400, "login.html", gin.H{"error": "Пользователь уже зарегистрирован!", "signupUser": signupUser})
		return
	}

	if err := signupUser.Validate(); err != nil {
		c.HTML(400, "login.html", gin.H{"error": err.Error(), "signupUser": signupUser})
		return
	}

	var errHash error
	signupUser.Password, errHash = utils.GenerateHashPassword(signupUser.Password)

	if errHash != nil {
		c.HTML(500, "login.html", gin.H{"error": "Ошибка: Генерация хэш-пароля!", "signupUser": signupUser})
		return
	}

	signupUser.Role = "user"

	user := signupUser.GetUser()

	models.DB.Create(&user)

	c.Redirect(302, "/auth?login")
}

func Logout(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", "localhost", false, true)
	fmt.Println("logout...")
	c.Redirect(http.StatusFound, "/auth?login")
}
