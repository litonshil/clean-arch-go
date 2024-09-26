package services

import (
	"clean-arch/client/logger"
	"clean-arch/internal/domain"
	"clean-arch/utils/consts"
	"fmt"
)

func TransactionRollback(
	txc *domain.TXClient,
	loggeruc logger.LogClient,
	entity consts.Entity,
	action consts.Action,
) error {
	if err := txc.Rollback(); err != nil {
		loggeruc.Error(
			fmt.Sprintf(
				"error occurred while transaction rollbacked for %v %v",
				entity,
				action,
			),
			err,
		)
		return err
	}

	loggeruc.Info("transaction rollbacked successfully ...")
	return nil
}

func TransactionCommit(
	txc *domain.TXClient,
	loggeruc logger.LogClient,
	entity consts.Entity,
	action consts.Action,
) error {
	if err := txc.Commit(); err != nil {
		loggeruc.Error(
			fmt.Sprintf(
				"error occurred while %v %v transaction commit",
				entity,
				action,
			),
			err,
		)
		return err
	}

	loggeruc.Info("transaction successfully committed...")
	return nil
}
