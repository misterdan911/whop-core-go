package controllers

import (
	"github.com/gin-gonic/gin"
	"whop-core-go/db"
	"whop-core-go/models"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"net/http"

	"os"
	"time"
)

type Auth struct {
}

type RequestBody struct {
	Email    string `form:"email" json:"email"`
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

type Response struct {
	Status  bool         `json:"status"`
	Message string       `json:"message"`
	Data    ResponseData `json:"data"`
}

func (auth Auth) Login(ctx *gin.Context) {

	var request RequestBody
	err := ctx.ShouldBind(&request)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	email := request.Email

	var users models.Users
	db.Whop.First(&users, "email = ?", email)

	isSame := isSame(request.Password, users.Password)

	var response Response

	if isSame {
		response.Status = true
		response.Message = "success"
	} else {
		response.Status = false
		response.Message = "failed"
	}

	response.Data.AccessToken, err = generateJWT()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, response)
}

/*
func Hash(str string) (string, error) {
    hashed, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
    return string(hashed), err
}
*/

func isSame(str string, hashed string) bool {
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
