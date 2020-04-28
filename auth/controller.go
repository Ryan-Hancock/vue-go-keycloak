package auth

import (
	"crypto/tls"
	"fmt"
	"net/http"

	"github.com/Nerzal/gocloak/v5"
	"github.com/gin-gonic/gin"
)

var (
	authService *AuthService
)

func Init() *AuthController {
	client := gocloak.NewClient("http://localhost:8080")
	restyClient := client.RestyClient()
	restyClient.SetDebug(true)
	restyClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})

	token, err := client.LoginAdmin("admin", "password", "master")
	if err != nil {
		panic(err)
	}

	return &AuthController{
		authService: &AuthService{
			Client:       client,
			Token:        token,
			ClientID:     "login-client",
			ClientSecret: "77b21a2d-3971-4ae7-9cb7-193625d7b2b3",
		},
	}
}

type AuthController struct {
	authService *AuthService
}

type CreateUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (a *AuthController) CreateUser(c *gin.Context) {
	token, err := a.authService.Client.LoginAdmin("admin", "password", "master")
	if err != nil {
		panic(err)
	}

	a.authService.Token = token

	var json CreateUserRequest
	if err = c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	uResp, err := a.authService.CreateUser(json)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": uResp})
	return
}

func (a AuthController) Login(c *gin.Context) {
	var json LoginRequest
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := a.authService.GetUserToken(json)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, token)
	return
}

func (a AuthController) Validate(c *gin.Context) {
	var json struct {
		AccessToken string `json:"accessToken"`
	}
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if ok, err := a.authService.RetrospectToken(json.AccessToken); !ok {
		if err != nil {
			fmt.Println(err)
		}
		c.JSON(http.StatusUnauthorized, gin.H{"Unauthorised": "token not valid"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "authorised"})

	return
}

func (a AuthController) Refresh(c *gin.Context) {
	var json struct {
		RefreshToken string `json:"refreshToken"`
	}
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := a.authService.RefreshToken(json.RefreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, token)

	return
}
