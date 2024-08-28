package model

import (
	"errors"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

type User struct {
	Uuid			string `json: id`
	UserName	string `json: username bson:"username, omitempty"`
	Email			string `json: email bson:"email, omitempty"`
	Password	string `json: password bson:"password, omitempty"`
	ApiUserID	string `json: apiUserID bson:"apiUserID, omitempty"`
	IsAdmin		bool   `json: isAdmin bson:"isAdmin"`
}

type JWTdesign struct{
	IsAdmin 	string
	Uuid			string
	ApiUserID	string
	jwt.RegisteredClaims
}

func (user *User) validate() (*User, error) {

	if strings.TrimSpace(user.UserName) == "" {
		return nil, errors.New("name required")
	}

	if strings.TrimSpace(user.Email) == "" {
		return nil, errors.New("mail required")
	}

	if strings.TrimSpace(user.Password) == "" {
		return nil, errors.New("password requied")
	}

	return user, nil

}
