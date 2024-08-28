package service

import (
	"errors"
	"example/model"
	"example/utils"
	"time"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
)
type RET struct {
	Token string
	Uuid string
}
	
// func LoginUser(user *model.User) (gin.H, error) {
func LoginUser(user *model.User) (*RET, error) {

	login := user.Email
	pass := user.Password

	result, error := model.LoginUser(login)
	// Check to exist User
	if error != nil {
		return nil, errors.New("user not found")
	}
	//[] user
	
	status := utils.CheckPassword(result.Password, pass)
	// Check hash password
	if !status {
		return nil, errors.New("invalid password")
	}
	//[] password
	/*
	• Generate JWT Token that will contain fields such as: is_admin,uuid,api_user_id
  • Set JWT Expiration to 15 minutes
	As Below |
	*/
	tokenStr := string(result.Uuid) + "||" + string(result.ApiUserID) + "||" + strconv.FormatBool(result.IsAdmin)
	loginToken := &model.JWTdesign{
		Uuid: result.Uuid,
		//IsAdmin: user.IsAdmin,
		//ApiUserID: user.ApiUserID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)), 
			Subject:   tokenStr,
		},
	}

	token, err := CreateJWTReg(loginToken)

	if err != nil {
		return nil, errors.New("JWT not generated")
		// return nil, errors.New("JWT not generated")
	}

	// return token, nil
	return &RET {
		Token: token,
		Uuid: result.Uuid,
	}, nil
}

func ValidReq(data *model.User) bool {
	if data.UserName == "" || data.Email == "" || data.Password == "" {
		return false
	}
	return true
}
