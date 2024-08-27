package service

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	"growth/conf"
	"growth/dbhelper"
	"growth/models"
	"log"
	"testing"
	"time"
)

func initDB() {
	time.Local = time.UTC
	conf.LoadConfigs()
	dbhelper.InitDb()
}

func TestCoinTaskService_Save(t *testing.T) {
	initDB()
	// 初始化一个Service对象
	s := NewCoinTaskService(context.Background())
	data := models.TbCoinTask{
		Id:    0,
		Task:  "post article",
		Coin:  10,
		Limit: 10,
	}
	if err := s.Save(&data); err != nil {
		t.Errorf("Save(%+v) error=%v", data, err)
	} else {
		log.Printf("Save data=%+v\n", data)
	}
}

func TestCoinTaskService_GetByTask(t *testing.T) {
	initDB()
	s := NewCoinTaskService(context.Background())
	task := "post article"
	if data, err := s.GetByTask(task); err != nil {
		t.Errorf("GetByTask(%s) error=%v", task, err)
	} else {
		log.Printf("GetByTask(%s) data=%v\n", task, data)
	}
}

func TestCoinTaskService_FindAll(t *testing.T) {
	initDB()
	s := NewCoinTaskService(context.Background())
	if dataList, err := s.FindAll(); err != nil {
		t.Errorf("FindAll() error=%v", err)
	} else {
		log.Printf("FindAll() data=%v\n", dataList)
	}
}

func TestGradeInfoService_Save(t *testing.T) {
	initDB()
	s := NewGradeInfoService(context.Background())
	data := models.TbGradeInfo{
		Id:          0,
		Title:       "初级",
		Description: "初级用户",
		Score:       0,
		Expired:     0,
	}
	if err := s.Save(&data); err != nil {
		t.Errorf("Save(%+v) error=%v", data, err)
	} else {
		log.Printf("Save data=%+v\n", data)
	}
}

func TestGradeInfoService_Get(t *testing.T) {
	initDB()
	s := NewGradeInfoService(context.Background())
	if data, err := s.Get(1); err != nil {
		t.Errorf("Get(1) error=%v", err)
	} else {
		log.Printf("Get(1) data=%+v\n", data)
	}
}

func TestGradeInfoService_FindAll(t *testing.T) {
	initDB()
	s := NewGradeInfoService(context.Background())
	if data, err := s.FindAll(); err != nil {
		t.Errorf("FindAll error=%v", err)
	} else {
		log.Printf("FindAll data=%+v\n", data)
	}
}
