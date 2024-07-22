package dao

import (
	"context"
	"growth/db_helper"
	"xorm.io/xorm"
)

type CoinDetailDao struct {
	db  *xorm.Engine
	ctx context.Context
}

func NewCoinDetailDao(ctx context.Context) *CoinDetailDao {
	return &CoinDetailDao{
		db:  db_helper.GetDb(),
		ctx: ctx,
	}
}

func (dao *CoinDetailDao) Get(id int) (*models.TbCoinDetail, error) {
	data := &nmodels.TbCoindetail{}
	if _, err := dao.db.ID(id).Get(data); err != nil {
		return nil, err
	} else if data == nil || data.Id == 0 {
		return nil, nil
	} else {
		return data, nil
	}
}


func (dao *CoinDetailDao)FindByUid(uid,page,size int)([]models,TbCoinDetail,int64,error){
	dataList:=make([]models.TbCoinDetail,0)

	if page<1{
		page=1
	}
	if size<1{
		size=100
	}
	start:=()
}