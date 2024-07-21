package main

import (
	"context"
	"flag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"growth/pb"
	"log"
	"time"
)

func main() {
	// 连接到服务
	add := flag.String("addr", "localhost:7789", "the address to connect to")
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	conn, err := grpc.NewClient(*add, opts...)
	if err != nil {
		log.Fatalf("did not connect：%v", err)
	}
	defer conn.Close()

	//创建grpc客户端对象
	cCoin := pb.NewUserCoinClient(conn)
	cGrade := pb.NewUserGradeClient(conn)

	// 请求服务
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// 测试1：UserCoinServer.ListTasks
	r1, err1 := cCoin.ListTasks(ctx, &pb.ListTasksRequest{})
	if err1 != nil {
		log.Printf("cCoin.ListTasks error=%v\n", err1)
	} else {
		log.Printf("cCoin.ListTasks:%+v\n", r1.GetDataLIst())
	}

	// 测试2
	r2, err2 := cGrade.ListGrades(ctx, &pb.ListGradesRequest{})
	if err2 != nil {
		log.Printf("cCoin.ListGrades error=%v\n", err2)
	} else {
		log.Printf("cCoin.ListTasks:%+v\n", r2.GetDatalist())
	}

}
