package models

import "time"

type TbCoinDetail struct {
	Id         int        `xorm:"not null pk autoincr INT(11)"`
	Uid        int        `xorm:"not null default 0 comment('用户id') index(uid) INT(11)"`
	TaskId     int        `xorm:"not null default - comment('任务id') index(uid) INT(11)"`
	Coin       int        `xorm:"not null default 0 comment('积分,正数是奖励，负数是惩罚') INT(11)"`
	SysCreated *time.Time `xorm:"not null deafult CURRENT_TIMESTAMP comment('创建时间') DATETIME"`
	SysUpdated *time.Time `xrom:"not null default CURRENT_TIMESTAMP comment('修改时间') DATETIME"`
}

type TbCoinTask struct {
	Id         int        `xorm:"not null pk autoincr INT(11)"`
	Task       string     `xorm:"not null default '' comment('任务名称，必须唯一') unique VARCHAR(255)"`
	Coin       int        `xorm:"not null default 0 comment('积分数，证书是奖励积分，负数是惩罚积分，0') INT(11)"`
	Limit      int        `xorm:"not null default 0 comment('每日限额，默认0不限制') INT(11)"`
	Start      *time.Time `xorm:"not null deafult CURRENT_TIMESTAMP comment('生效开始时间') DATETIME"`
	SysCreated *time.Time `xorm:"not null deafult CURRENT_TIMESTAMP comment('创建时间') DATETIME"`
	SysUpdated *time.Time `xrom:"not null default CURRENT_TIMESTAMP comment('修改时间') DATETIME"`
	Status     int        `xorm:"not null default 0 comment('状态，默认0整数，1删除') INT(11)"`
}
