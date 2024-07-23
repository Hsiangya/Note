package dao

import (
	"context"
	"growth/comm"
	"growth/dbhelper"
	"growth/models"

	"xorm.io/xorm"
)

// CoinUserDao process table tb_km_article, model TbCoinUser
type CoinUserDao struct {
	db  *xorm.Engine
	ctx context.Context
}

// NewCoinUserDao initialize CoinUserDao instance.
func NewCoinUserDao(ctx context.Context) *CoinUserDao {
	return &CoinUserDao{
		db:  dbhelper.GetDb(),
		ctx: ctx,
	}
}

// Get model by id.
func (dao *CoinUserDao) Get(id int) (*models.TbCoinUser, error) {
	data := &models.TbCoinUser{}
	if _, err := dao.db.ID(id).Get(data); err != nil {
		return nil, err
	} else if data == nil || data.Id == 0 {
		return nil, nil
	} else {
		return data, nil
	}
}

// GetByUid get model by uid
func (dao *CoinUserDao) GetByUid(uid int) (*models.TbCoinUser, error) {
	data := models.TbCoinUser{}
	sess := dao.db.Where("`uid`=?", uid)
	_, err := sess.Get(&data)
	return &data, err
}

// FindAllPager get all models
func (dao *CoinUserDao) FindAllPager(page, size int) ([]models.TbCoinUser, int64, error) {
	datalist := make([]models.TbCoinUser, 0)
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
func (dao *CoinUserDao) Insert(data *models.TbCoinUser) error {
	data.SysCreated = comm.Now()
	data.SysUpdated = comm.Now()
	_, err := dao.db.Insert(data)
	return err
}

// Update one row
func (dao *CoinUserDao) Update(data *models.TbCoinUser, musColumns ...string) error {
	sess := dao.db.ID(data.Id)
	if len(musColumns) > 0 {
		sess.MustCols(musColumns...)
	}
	_, err := sess.Update(data)
	return err
}

// Save with Insert and Update
func (dao *CoinUserDao) Save(data *models.TbCoinUser, musColumns ...string) error {
	if data.Id > 0 {
		return dao.Update(data, musColumns...)
	} else {
		return dao.Insert(data)
	}
}
