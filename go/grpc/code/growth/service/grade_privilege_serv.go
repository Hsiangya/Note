package service

import (
	"context"
	"growth/dao"
	"growth/models"
)

// GradePrivilegeService service for knowledge article
type GradePrivilegeService struct {
	cxt               context.Context
	daoGradePrivilege *dao.GradePrivilegeDao
}

// NewGradePrivilegeService new instance of GradePrivilegeService
func NewGradePrivilegeService(ctx context.Context) *GradePrivilegeService {
	return &GradePrivilegeService{
		cxt:               ctx,
		daoGradePrivilege: dao.NewGradePrivilegeDao(ctx),
	}
}

// Get model by id.
func (s *GradePrivilegeService) Get(id int) (*models.TbGradePrivilege, error) {
	return s.daoGradePrivilege.Get(id)
}

// FindByGrade get model by gradeId
func (s *GradePrivilegeService) FindByGrade(gradeId int) ([]models.TbGradePrivilege, error) {
	return s.daoGradePrivilege.FindByGrade(gradeId)
}

// FindAll get all models
func (s *GradePrivilegeService) FindAll() ([]models.TbGradePrivilege, error) {
	return s.daoGradePrivilege.FindAll()
}

// Save with Insert and Update
func (s *GradePrivilegeService) Save(data *models.TbGradePrivilege, musColumns ...string) error {
	return s.daoGradePrivilege.Save(data, musColumns...)
}
