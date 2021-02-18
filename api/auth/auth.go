package auth

import (
	"errors"
	"fmt"
	"new/test/project/api/constants"
	"new/test/project/api/model"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWT struct {
	expiry int64
}

func New() *JWT {
	return &JWT{expiry: 72}
}

type MyProjectClaims struct {
	ID int `json:"id"`
	jwt.StandardClaims
}

func (j *JWT) CreateToken(user *model.User) (string, error) {

	claims := MyProjectClaims{
		int(user.ID),
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(j.expiry)).Unix(),
			Issuer:    constants.JWTIssuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(constants.Secret))
}

func (j *JWT) ValidateToken(token string) (bool, error) {
	err := at(time.Unix(0, 0), func() error {
		jtoken, err := jwt.ParseWithClaims(token, &MyProjectClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(constants.Secret), nil
		})
		if err != nil {
			return err
		}

		if claims, ok := jtoken.Claims.(*MyProjectClaims); ok && jtoken.Valid {
			fmt.Printf("%v %v", claims.ID, claims.StandardClaims.ExpiresAt)
			return nil
		}

		return errors.New("Invalid token")
	})
	if err != nil {
		return false, err
	}

	return true, nil
}

func at(t time.Time, f func() error) error {
	jwt.TimeFunc = func() time.Time {
		return t
	}

	err := f()
	if err != nil {
		return err
	}

	jwt.TimeFunc = time.Now
	return nil
}
