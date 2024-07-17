package main

import (
	"google.golang.org/grpc"
	"growth/pb"
	"growth/ugserver"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", "80")
	if err != nil {
		log.Fatalf("failed to listen: %s", err.Error())
	}

	// 创建grpc服务
	s := grpc.NewServer()

	// 注册服务
	pb.RegisterUserCoinServer(s, &ugserver.UgCoinServer{})
	pb.RegisterUserGradeServer(s, &ugserver.UgGradeServer{})

	// 启动服务
	log.Printf("server listening at %v\n", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}
