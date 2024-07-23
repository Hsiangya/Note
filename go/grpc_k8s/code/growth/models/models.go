package models

import (
	"time"
)

type TbCoinDetail struct {
	Id         uint      `xorm:"not null pk autoincr UNSIGNED INT"`
	Uid        int       `xorm:"not null default 0 comment('用户id') index(uid) INT"`
	TaskId     int       `xorm:"not null default 0 comment('任务id') index(uid) INT"`
	Coin       int       `xorm:"not null default 0 comment('积分，正数是奖励，负数是惩罚') INT"`
	SysCreated time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') DATETIME"`
	SysUpdated time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('修改时间') DATETIME"`
}

type TbCoinTask struct {
	Id         uint      `xorm:"not null pk autoincr UNSIGNED INT"`
	Task       string    `xorm:"not null default '' comment('任务名称，必须唯一') unique VARCHAR(255)"`
	Coin       int       `xorm:"not null default 0 comment('积分数，正数是奖励积分，负数是惩罚积分，0需要外部调用传值') INT"`
	Limit      int       `xorm:"not null default 0 comment('每日限额，默认0不限制') INT"`
	Start      time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('生效开始时间') DATETIME"`
	SysCreated time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') DATETIME"`
	SysUpdated time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('修改时间') DATETIME"`
	SysStatus  int       `xorm:"not null default 0 comment('状态，默认0整除，1删除') INT"`
}

type TbCoinUser struct {
	Id         uint      `xorm:"not null pk autoincr UNSIGNED INT"`
	Uid        int       `xorm:"not null default 0 comment('用户id') unique INT"`
	Coins      int       `xorm:"not null default 0 comment('总积分') INT"`
	SysCreated time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') DATETIME"`
	SysUpdated time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('修改时间') DATETIME"`
}

type TbGradeInfo struct {
	Id          uint      `xorm:"not null pk autoincr UNSIGNED INT"`
	Title       string    `xorm:"not null comment('等级名称') VARCHAR(255)"`
	Description string    `xorm:"not null comment('等级描述信息') VARCHAR(3000)"`
	Score       int       `xorm:"not null default 0 comment('等级最高的成长数值') INT"`
	Expired     int       `xorm:"not null default 0 comment('有效期，单位:天，默认0永不过期') INT"`
	SysCreated  time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') DATETIME"`
	SysUpdated  time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('修改时间') DATETIME"`
}

type TbGradePrivilege struct {
	Id          uint      `xorm:"not null pk autoincr UNSIGNED INT"`
	GradeId     int       `xorm:"not null default 0 comment('等级id') index INT"`
	Product     string    `xorm:"not null default '' comment('产品') VARCHAR(255)"`
	Function    string    `xorm:"not null default '' comment('功能') VARCHAR(255)"`
	Description string    `xorm:"not null default '' comment('描述信息') VARCHAR(3000)"`
	Expired     int       `xorm:"not null default 0 comment('有效期，单位:天，默认0永不过期') INT"`
	SysCreated  time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') DATETIME"`
	SysUpdated  time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('修改时间') DATETIME"`
	SysStatus   int       `xorm:"not null default 0 comment('状态，默认0整除，1删除') INT"`
}

type TbGradeUser struct {
	Id         uint      `xorm:"not null pk autoincr UNSIGNED INT"`
	Uid        int       `xorm:"not null default 0 comment('用户id') index INT"`
	GradeId    int       `xorm:"not null default 0 comment('等级id') INT"`
	Expired    time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('过期时间') DATETIME"`
	Score      int       `xorm:"not null default 0 comment('成长数值') INT"`
	SysCreated time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') DATETIME"`
	SysUpdated time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('修改时间') DATETIME"`
}
