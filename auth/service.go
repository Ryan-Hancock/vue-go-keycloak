package auth

import (
	"fmt"

	"github.com/Nerzal/gocloak/v5"
)

var ()

type AuthService struct {
	Client       gocloak.GoCloak
	Token        *gocloak.JWT
	ClientID     string
	ClientSecret string
}

// CreateUser...
func (as *AuthService) CreateUser(cur CreateUserRequest) (string, error) {
	user := gocloak.User{
		Email:    &cur.Email,
		Enabled:  gocloak.BoolP(true),
		Username: &cur.Username,
	}

	fmt.Println(as.Token)
	uResp, err := as.Client.CreateUser(as.Token.AccessToken, "master", user)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	err = as.Client.SetPassword(
		as.Token.AccessToken,
		uResp,
		"master",
		cur.Password,
		false,
	)

	return uResp, err
}

func (as *AuthService) GetUserToken(lr LoginRequest) (*gocloak.JWT, error) {
	token, err := as.Client.Login(
		as.ClientID,
		as.ClientSecret,
		"master",
		lr.Username,
		lr.Password,
	)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return token, err
}

func (as *AuthService) RetrospectToken(token string) (bool, error) {
	res, err := as.Client.RetrospectToken(
		token,
		as.ClientID,
		as.ClientSecret,
		"master",
	)
	if err != nil {
		fmt.Println(err)
		return false, err
	}

	return *res.Active, err
}

func (as *AuthService) RefreshToken(refreshToken string) (*gocloak.JWT, error) {
	token, err := as.Client.RefreshToken(
		refreshToken,
		as.ClientID,
		as.ClientSecret,
		"master",
	)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return token, err
}
