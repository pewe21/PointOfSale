package authentication

import (
	"context"
	"errors"
	"github.com/pewe21/PointOfSale/dto"
	"github.com/pewe21/PointOfSale/internal/config"
	"github.com/pewe21/PointOfSale/internal/domain"
	"github.com/pewe21/PointOfSale/internal/util"
	"golang.org/x/crypto/bcrypt"
)

type service struct {
	domain.UserRepository
	*config.Jwt
}

func NewService(userRepository domain.UserRepository, Jwt *config.Jwt) domain.AuthService {
	return &service{UserRepository: userRepository, Jwt: Jwt}
}

func (s service) SignIn(ctx context.Context, req dto.SignInRequest) (string, error) {
	user, err := s.UserRepository.FindByEmail(ctx, req.Username)
	if err != nil {
		return "", errors.New("username or password incorrect")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return "", errors.New("username or password incorrect")
	}

	jwt := util.NewJwtUtils(s.Jwt)

	t, err := jwt.CreateToken(user)

	if err != nil {
		return "", err
	}

	return t, nil
}
