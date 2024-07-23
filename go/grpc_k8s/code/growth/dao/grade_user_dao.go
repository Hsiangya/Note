package dao

import (
	"context"
	"growth/comm"
	"growth/dbhelper"
	"growth/models"

	"xorm.io/xorm"
)

// GradeUserDao process table tb_km_article, model TbGradeUser
type GradeUserDao struct {
	db  *xorm.Engine
	ctx context.Context
}

// NewGradeUserDao initialize GradeUserDao instance.
func NewGradeUserDao(ctx context.Context) *GradeUserDao {
	return &GradeUserDao{
		db:  dbhelper.GetDb(),
		ctx: ctx,
	}
}

// Get model by id.
func (dao *GradeUserDao) Get(id int) (*models.TbGradeUser, error) {
	data := &models.TbGradeUser{}
	if _, err := dao.db.ID(id).Get(data); err != nil {
		return nil, err
	} else if data == nil || data.Id == 0 {
		return nil, nil
	} else {
		return data, nil
	}
}

// GetByUid get model by uid
func (dao *GradeUserDao) GetByUid(uid int) (*models.TbGradeUser, error) {
	data := models.TbGradeUser{}
	sess := dao.db.Where("`uid`=?", uid)
	_, err := sess.Get(&data)
	return &data, err
}

// FindAllPager get all models
func (dao *GradeUserDao) FindAllPager(page, size int) ([]models.TbGradeUser, int64, error) {
	datalist := make([]models.TbGradeUser, 0)
	if page < 1 {
		page = 1
	}
	if size < 1 {
		size = 100
	}
	start := (page - 1) * size
	total, err := dao.db.Desc("id").Limit(size, start).FindAndCount(&datalist)
	return datalist, total, err
}

// Insert one row
func (dao *GradeUserDao) Insert(data *models.TbGradeUser) error {
	data.SysCreated = comm.Now()
	data.SysUpdated = comm.Now()
	_, err := dao.db.Insert(data)
	return err
}

// Update one row
func (dao *GradeUserDao) Update(data *models.TbGradeUser, musColumns ...string) error {
	sess := dao.db.ID(data.Id)
	if len(musColumns) > 0 {
		sess.MustCols(musColumns...)
	}
	_, err := sess.Update(data)
	return err
}

// Save with Insert and Update
func (dao *GradeUserDao) Save(data *models.TbGradeUser, musColumns ...string) error {
	if data.Id > 0 {
		return dao.Update(data, musColumns...)
	} else {
		return dao.Insert(data)
	}
}
