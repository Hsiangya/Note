package dao

import (
	"context"
	"growth/comm"
	"growth/dbhelper"
	"growth/models"
	"xorm.io/xorm"
)

// GradeInfoDao process table tb_km_article, model TbGradeInfo
type GradeInfoDao struct {
	db  *xorm.Engine
	ctx context.Context
}

// NewGradeInfoDao initialize GradeInfoDao instance.
func NewGradeInfoDao(ctx context.Context) *GradeInfoDao {
	return &GradeInfoDao{
		db:  dbhelper.GetDb(),
		ctx: ctx,
	}
}

// Get model by id.
func (dao *GradeInfoDao) Get(id int) (*models.TbGradeInfo, error) {
	data := &models.TbGradeInfo{}
	if _, err := dao.db.ID(id).Get(data); err != nil {
		return nil, err
	} else if data == nil || data.Id == 0 {
		return nil, nil
	} else {
		return data, nil
	}
}

// FindAll get all models
func (dao *GradeInfoDao) FindAll() ([]models.TbGradeInfo, error) {
	datalist := make([]models.TbGradeInfo, 0)
	err := dao.db.Asc("score").Find(&datalist)
	return datalist, err
}

// Insert one row
func (dao *GradeInfoDao) Insert(data *models.TbGradeInfo) error {
	data.SysCreated = comm.Now()
	data.SysUpdated = comm.Now()
	_, err := dao.db.Insert(data)
	return err
}

// Update one row
func (dao *GradeInfoDao) Update(data *models.TbGradeInfo, musColumns ...string) error {
	sess := dao.db.ID(data.Id)
	if len(musColumns) > 0 {
		sess.MustCols(musColumns...)
	}
	_, err := sess.Update(data)
	return err
}

// Save with Insert and Update
func (dao *GradeInfoDao) Save(data *models.TbGradeInfo, musColumns ...string) error {
	if data.Id > 0 {
		return dao.Update(data, musColumns...)
	} else {
		return dao.Insert(data)
	}
}
