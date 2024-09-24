package domain

import (
	"clean-arch/types"
	"context"
)

type AuthUseCase interface {
	Login(ctx context.Context, req *types.LoginReq) (*types.LoginResp, error)
}

type AuthRepo interface {
	Login(req *types.LoginReq) (*types.LoginResp, error)
}
