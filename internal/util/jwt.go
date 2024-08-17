package util

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/pewe21/PointOfSale/internal/config"
	"github.com/pewe21/PointOfSale/internal/domain"
	"time"
)

type JwtUtils struct {
	cnf *config.Jwt
}

func NewJwtUtils(jwt *config.Jwt) *JwtUtils {
	return &JwtUtils{cnf: jwt}
}

func (j *JwtUtils) CreateToken(user domain.User) (string, error) {

	claims := jwt.MapClaims{
		"id":    user.Id,
		"name":  user.Name,
		"email": user.Email,
		"exp":   time.Now().Add(time.Hour * time.Duration(j.cnf.Exp)).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(j.cnf.Secret))
	if err != nil {
		return "", errors.New("token sign failed")
	}
	return t, nil
}
