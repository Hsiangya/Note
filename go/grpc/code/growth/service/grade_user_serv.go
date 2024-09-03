package service

import (
	"context"
	"growth/dao"
	"growth/models"
)

// GradeUserService service for knowledge article
type GradeUserService struct {
	cxt          context.Context
	daoGradeUser *dao.GradeUserDao
}

// NewGradeUserService new instance of GradeUserService
func NewGradeUserService(ctx context.Context) *GradeUserService {
	return &GradeUserService{
		cxt:          ctx,
		daoGradeUser: dao.NewGradeUserDao(ctx),
	}
}

// Get model by id.
func (s *GradeUserService) Get(id int) (*models.TbGradeUser, error) {
	return s.daoGradeUser.Get(id)
}

// GetByUid get model by uid
func (s *GradeUserService) GetByUid(uid int) (*models.TbGradeUser, error) {
	return s.daoGradeUser.GetByUid(uid)
}

// FindAllPager get all models
func (s *GradeUserService) FindAllPager(page, size int) ([]models.TbGradeUser, int64, error) {
	return s.daoGradeUser.FindAllPager(page, size)
}

// Save with Insert and Update
func (s *GradeUserService) Save(data *models.TbGradeUser, musColumns ...string) error {
	return s.daoGradeUser.Save(data, musColumns...)
}
