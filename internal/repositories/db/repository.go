package db

import (
	"clean-arch/infra/logger"
	"clean-arch/internal/domain"
	"clean-arch/internal/services/transaction"
	"context"

	"gorm.io/gorm"
)

type Repository struct {
	client *gorm.DB
	lc     *logger.LogClient
}

func NewRepository(client *gorm.DB, lc *logger.LogClient) *Repository {
	return &Repository{
		client: client,
		lc:     lc,
	}
}

func (repo *Repository) CreateTransaction(ctx context.Context) (*domain.TXClient, error) {
	tx := repo.client.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	repo.lc.Info("transaction started...")
	return transaction.NewTXClient(ctx, tx), nil
}

func (repo *Repository) dbClient(txc *domain.TXClient) *gorm.DB {
	if txc == nil {
		return repo.client
	}

	return txc.Get().(*gorm.DB)
}
