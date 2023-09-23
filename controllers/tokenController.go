package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// logging and authorization
var jwtKey = []byte("secret_key")

// Store & access user information
var users = map[string]string{
	"user1": "password1",
	"user2": "password2",
}

// Credentials creates a struct with two fields, user and pass, which are both strings. A struct is a custom data type that can be used to store related data.
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Claims struct which contains a field named Username of type string and a field of type jwt.StandardClaims.
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// Login function takes a username and password and checks if the credentials are valid.
func Login(c *gin.Context) {
	// decode JSON data sent in an HTTP request.
	// var "credentials" is of type "Credentials".json.NewDecoder() is used to create a new decoder object which will be used to decode the JSON data sent in the request.
	//The decoded data is then stored in the "credentials" variable. If there is an error while decoding the data, the http status code "400" is sent to the client and the code execution is stopped.
	var credentials Credentials
	err := json.NewDecoder(c.Request.Body).Decode(&credentials)
	if err != nil {
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}

	expectedPassword, ok := users[credentials.Username]

	if !ok || expectedPassword != credentials.Password {
		c.Writer.WriteHeader(http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().Add(time.Minute * 5)

	claims := &Claims{
		Username: credentials.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(c.Writer,
		&http.Cookie{
			Name:    "token",
			Value:   tokenString,
			Expires: expirationTime,
		})

}

// type TokenRequest struct {
// 	Email    string `json:"email"`
// 	Password string `json:"password"`
// }

// func GenerateToken(context *gin.Context) {
// 	var request TokenRequest
// 	var user models.User
// 	if err := context.ShouldBindJSON(&request); err != nil {
// 		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		context.Abort()
// 		return
// 	}

// 	// check if email exists and password is correct
// 	record := config.Database.Where("email = ?", request.Email).First(&user)
// 	if record.Error != nil {
// 		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
// 		context.Abort()
// 		return
// 	}

// 	credentialError := user.CheckPassword(request.Password)
// 	if credentialError != nil {
// 		context.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
// 		context.Abort()
// 		return
// 	}

// 	tokenString, err := auth.GenerateJWT(user.Email, user.Username)
// 	if err != nil {
// 		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		context.Abort()
// 		return
// 	}
// 	context.JSON(http.StatusOK, gin.H{"token": tokenString})
// }
