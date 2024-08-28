package service

import (
	"errors"
	"example/model"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
)

var User interface{}

func CreateJWTReg(claims *model.JWTdesign) (string, error) {

	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims.RegisteredClaims).SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		return "", err
	}

	return token, nil
}

func ValidateJWT(tokens string) (interface{}, error) {
	token, _ := jwt.Parse(tokens, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			fmt.Println("nil comming")
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	claims, ok := token.Claims.(jwt.MapClaims)

	fmt.Println("claim status OK: ",token)
	fmt.Println("token valid: ",token.Valid)
	fmt.Println("expiryTime: ",claims["exp"].(float64))
	if ok && token.Valid {
		User = claims["sub"]
	} else if float64(time.Now().Unix()) > claims["exp"].(float64) && token.Valid {
		return nil, errors.New("JWT expiry")
	} else {
		return nil, errors.New("invalid token")
	}

	fmt.Println("user mail: ", ok)

	return User, nil
}

func Uuid(s string) string {
	return s
}
