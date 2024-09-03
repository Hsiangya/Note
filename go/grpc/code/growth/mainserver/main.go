package main

import (
	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"growth/conf"
	"growth/dbhelper"
	"growth/pb"
	"growth/ugserver"
	"log"
	"net"
	"time"
)

func initDb() {
	// default UTC time location
	time.Local = time.UTC
	// Load global config
	conf.LoadConfigs()
	// Initialize db
	dbhelper.InitDb()
}
func main() {
	// 初始化数据库实例
	initDb()

	lis, err := net.Listen("tcp", ":7789")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	// 注册服务
	pb.RegisterUserCoinServer(s, &ugserver.UgCoinServer{})
	pb.RegisterUserGradeServer(s, &ugserver.UgGradeServer{})
	reflection.Register(s)
	// 启动服务
	log.Printf("server listening at %v\n", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
