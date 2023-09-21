package auth

import (
	"errors"
	"math/rand"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	authRepo AuthRepository
}

func (s *AuthService) SignUp(p *SignupInput) error {
	password := []byte(p.Password)
	hashed, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	p.Password = string(hashed)
	err = s.authRepo.Create(p)
	if err != nil {
		return err
	}

	return nil
}

func (s *AuthService) Signin(p *SigninInput) (*string, error) {
	user, err := s.authRepo.FindByUsername(&p.Username)
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(p.Password))
	if err != nil {
		return nil, errors.New("password didn't match")
	}

	token := s.GenerateToken()
	err = s.authRepo.CreateToken(&user.ID, &token)
	if err != nil {
		return nil, err
	}
	return &token, nil
}

func (s *AuthService) Signout(p *SignoutInput) error {
	err := s.authRepo.SoftDeleteToken(p)
	if err != nil {
		return err
	}

	return nil
}

func (s *AuthService) GenerateToken() string {
	token := "token"
	str := "abcdefghijklmnopqrstuvwxyz1234567890"
	shuff := []rune(str)
	rand.Shuffle(len(shuff), func(i, j int) {
		shuff[i], shuff[j] = shuff[j], shuff[i]
	})

	token += string(shuff)
	return token
}
