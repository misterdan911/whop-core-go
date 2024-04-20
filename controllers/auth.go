package controllers

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"whop-core-go/db"
	"whop-core-go/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"net/http"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"

	"os"
	"time"
)

type Auth struct {
}

type RequestBody struct {
	Email    string `form:"email" json:"email" binding:"required,email"`
	Password string `form:"password" json:"password"`
}

type ResponseData struct {
	// ID          int    `json:"id"`
	// FullName    string `json:"fullName"`
	// UserName    string `json:"username"`
	// Avatar      string `json:"avatar"`
	// Email       string `json:"avatar"`
	// Role        string `json:"avatar"`
	AccessToken string `json:"accessToken"`
}

type ResponseSuccess struct {
	Status  bool         `json:"status"`
	Message string       `json:"message"`
	Data    ResponseData `json:"data"`
}

type ResponseFailed struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

type ErrorMsg struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func getErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "lte":
		return "Should be less than " + fe.Param()
	case "gte":
		return "Should be greater than " + fe.Param()
	}
	return "Unknown error"
}

func (auth Auth) Login(ctx *gin.Context) {

	var request RequestBody
	var isValidUser bool = true
	var jwtToken string

	err := ctx.ShouldBind(&request)

	if err != nil {

		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]ErrorMsg, len(ve))
			for i, fe := range ve {
				out[i] = ErrorMsg{fe.Field(), getErrorMsg(fe)}
			}
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": out})
		}

		return
	}

	email := request.Email

	var users models.Users
	err = db.Whop.First(&users, "email = ?", email).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		isValidUser = false
	}

	passwordIsSame := passwordIsSame(request.Password, users.Password)

	if !passwordIsSame {
		isValidUser = false
	}

	jwtToken, err = generateJWT()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if isValidUser {
		var response ResponseSuccess

		response.Status = true
		response.Message = "success"
		response.Data.AccessToken = jwtToken

		ctx.JSON(200, response)
	} else {
		//var response ResponseFailed
		//
		//response.Status = false
		//response.Message = "Username or password isn't correct."

		response := gin.H{
			"status":  false,
			"message": "Username or password isn't correct.",
		}

		ctx.JSON(200, response)
	}

}

func getSimpleError() {

}

/*
func Hash(str string) (string, error) {
    hashed, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
    return string(hashed), err
}
*/

func passwordIsSame(str string, hashed string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(str)) == nil
}

/*
func generateJWT() (string, error) {

	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":        time.Now().Add(10 * time.Minute),
		"authorized": true,
		"user":       "username",
	})

	var jwtSecret = []byte(os.Getenv("JWT_SECRET"))
	tokenString, err := token.SignedString(jwtSecret)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
*/

/*
func generateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(10 * time.Minute)
	claims["authorized"] = true
	claims["user"] = "username"

	var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "Signing Error", err
	}

	return tokenString, nil
}
*/

func generateJWT() (string, error) {

	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo": "bar",
		"nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	var jwtSecret = []byte(os.Getenv("JWT_SECRET"))
	tokenString, err := token.SignedString(jwtSecret)

	return tokenString, err
}
