package service

import (
	"context"
	"growth/dao"
	"growth/models"
)

// CoinUserService service for knowledge article
type CoinUserService struct {
	cxt         context.Context
	daoCoinUser *dao.CoinUserDao
}

// NewCoinUserService new instance of CoinUserService
func NewCoinUserService(ctx context.Context) *CoinUserService {
	return &CoinUserService{
		cxt:         ctx,
		daoCoinUser: dao.NewCoinUserDao(ctx),
	}
}

// Get model by id.
func (s *CoinUserService) Get(id int) (*models.TbCoinUser, error) {
	return s.daoCoinUser.Get(id)
}

// GetByUid get model by uid
func (s *CoinUserService) GetByUid(uid int) (*models.TbCoinUser, error) {
	return s.daoCoinUser.GetByUid(uid)
}

// FindAllPager get all models
func (s *CoinUserService) FindAllPager(page, size int) ([]models.TbCoinUser, int64, error) {
	return s.daoCoinUser.FindAllPager(page, size)
}

// Save with Insert and Update
func (s *CoinUserService) Save(data *models.TbCoinUser, musColumns ...string) error {
	return s.daoCoinUser.Save(data, musColumns...)
}
