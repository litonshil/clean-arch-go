package db

import (
	"clean-arch/types"
)

func (repo *Repository) Login(user *types.LoginReq) (*types.LoginResp, error) {
	//if err := repo.db.Create(user).Error; err != nil {
	//	return err
	//}

	return nil, nil
}
