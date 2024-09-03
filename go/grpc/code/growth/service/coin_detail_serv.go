package service

import (
	"context"
	"growth/dao"
	"growth/models"
)

// CoinDetailService service for knowledge article
type CoinDetailService struct {
	cxt           context.Context
	daoCoinDetail *dao.CoinDetailDao
}

// NewCoinDetailService new instance of CoinDetailService
func NewCoinDetailService(ctx context.Context) *CoinDetailService {
	return &CoinDetailService{
		cxt:           ctx,
		daoCoinDetail: dao.NewCoinDetailDao(ctx),
	}
}

// Get model by id.
func (s *CoinDetailService) Get(id int) (*models.TbCoinDetail, error) {
	return s.daoCoinDetail.Get(id)
}

// FindByUid get models by uid
func (s *CoinDetailService) FindByUid(uid, page, size int) ([]models.TbCoinDetail, int64, error) {
	return s.daoCoinDetail.FindByUid(uid, page, size)
}

// FindAllPager get all models
func (s *CoinDetailService) FindAllPager(page, size int) ([]models.TbCoinDetail, int64, error) {
	return s.daoCoinDetail.FindAllPager(page, size)
}

// Save with Insert and Update
func (s *CoinDetailService) Save(data *models.TbCoinDetail, musColumns ...string) error {
	return s.daoCoinDetail.Save(data, musColumns...)
}
