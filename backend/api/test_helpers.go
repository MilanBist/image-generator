package api

import (
	"github.com/image-generator/internal/models"
)

type FakeUserStore struct {
	User models.Login
	ID   int
	Err  error
}

func (f *FakeUserStore) LoginUser(credentials models.Login) (int, error) {
	f.User = credentials
	return f.ID, f.Err
}

type FakeTokenService struct {
	AccessToken  string
	RefreshToken string
	Err          error
}

func (f *FakeTokenService) GenerateTokens(userId int64, email, requirement string) (string, string, error) {
	return f.RefreshToken, f.AccessToken, f.Err
}
