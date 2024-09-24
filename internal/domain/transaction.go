package domain

import (
	"context"
	"gorm.io/gorm"
)

type (
	TXClient struct {
		Ctx    context.Context
		Client *gorm.DB
	}

	TXRepo interface {
		CreateTransaction(ctx context.Context) (*TXClient, error)
	}

	TXUseCase interface {
		CreateTransaction(ctx context.Context) (*TXClient, error)
	}
)

func (txc *TXClient) Get() interface{} {
	return txc.Client
}

func (txc *TXClient) GetCtx() context.Context {
	return txc.Ctx
}

func (txc *TXClient) Rollback() error {
	return txc.Client.Rollback().Error
}

func (txc *TXClient) Commit() error {
	return txc.Client.Commit().Error
}
