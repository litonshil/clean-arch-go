package auth

import (
	"clean-arch/internal/domain"
	"clean-arch/types"
	"context"
)

type AuthService struct {
	repo domain.AuthRepo
}

func NewAuthService(authRepo domain.AuthRepo) *AuthService {
	return &AuthService{
		repo: authRepo,
	}
}

func (service *AuthService) Login(ctx context.Context, req *types.LoginReq) (*types.LoginResp, error) {
	_ = req
	_ = ctx
	//if err := service.repo.SignUp(user); err != nil {
	//	return errors.New("User was not created")
	//}

	return nil, nil
}
