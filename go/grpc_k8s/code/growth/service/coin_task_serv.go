package service

import (
	"context"
	"growth/dao"
	"growth/models"
)

// CoinTaskService service for knowledge article
type CoinTaskService struct {
	cxt         context.Context
	daoCoinTask *dao.CoinTaskDao
}

// NewCoinTaskService new instance of CoinTaskService
func NewCoinTaskService(ctx context.Context) *CoinTaskService {
	return &CoinTaskService{
		cxt:         ctx,
		daoCoinTask: dao.NewCoinTaskDao(ctx),
	}
}

// Get model by id.
func (s *CoinTaskService) Get(id int) (*models.TbCoinTask, error) {
	return s.daoCoinTask.Get(id)
}

// GetByTask get models by task
func (s *CoinTaskService) GetByTask(task string) (*models.TbCoinTask, error) {
	return s.daoCoinTask.GetByTask(task)
}

// FindAll get all models
func (s *CoinTaskService) FindAll() ([]models.TbCoinTask, error) {
	return s.daoCoinTask.FindAll()
}

// Save with Insert and Update
func (s *CoinTaskService) Save(data *models.TbCoinTask, musColumns ...string) error {
	return s.daoCoinTask.Save(data, musColumns...)
}
