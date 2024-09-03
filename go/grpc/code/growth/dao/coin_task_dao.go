package dao

import (
	"context"
	"growth/comm"
	"growth/dbhelper"
	"growth/models"
	"xorm.io/xorm"
)

// CoinTaskDao process table tb_km_article, model TbCoinTask
type CoinTaskDao struct {
	db  *xorm.Engine
	ctx context.Context
}

// NewCoinTaskDao initialize CoinTaskDao instance.
func NewCoinTaskDao(ctx context.Context) *CoinTaskDao {
	return &CoinTaskDao{
		db:  dbhelper.GetDb(),
		ctx: ctx,
	}
}

// Get model by id.
func (dao *CoinTaskDao) Get(id int) (*models.TbCoinTask, error) {
	data := &models.TbCoinTask{}
	if _, err := dao.db.ID(id).Get(data); err != nil {
		return nil, err
	} else if data == nil || data.Id == 0 {
		return nil, nil
	} else {
		return data, nil
	}
}

// GetByTask model by id.
func (dao *CoinTaskDao) GetByTask(task string) (*models.TbCoinTask, error) {
	data := &models.TbCoinTask{}
	if _, err := dao.db.Where("`task`=?", task).Get(data); err != nil {
		return nil, err
	} else if data == nil || data.Id == 0 {
		return nil, nil
	} else {
		return data, nil
	}
}

// FindAll get all models
func (dao *CoinTaskDao) FindAll() ([]models.TbCoinTask, error) {
	datalist := make([]models.TbCoinTask, 0)
	err := dao.db.Desc("id").Find(&datalist)
	return datalist, err
}

// Insert one row
func (dao *CoinTaskDao) Insert(data *models.TbCoinTask) error {
	data.SysCreated = comm.Now()
	data.SysUpdated = comm.Now()
	_, err := dao.db.Insert(data)
	return err
}

// Update one row
func (dao *CoinTaskDao) Update(data *models.TbCoinTask, musColumns ...string) error {
	sess := dao.db.ID(data.Id)
	if len(musColumns) > 0 {
		sess.MustCols(musColumns...)
	}
	_, err := sess.Update(data)
	return err
}

// Save with Insert and Update
func (dao *CoinTaskDao) Save(data *models.TbCoinTask, musColumns ...string) error {
	now := comm.Now()
	if data.Id > 0 {
		return dao.Update(data, musColumns...)
	} else {
		if data.Start == nil {
			data.Start = now
		}
		return dao.Insert(data)
	}
}
