package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"snowflake/dao/config"
	"snowflake/dao/logger"
	"snowflake/dao/register"
	"snowflake/handler"
	"snowflake/proto"
	"syscall"
)

func main() {
	// init config
	var cfn string // 从命令行获取可能的conf路径	goods_service -conf="./conf/config_qa.yaml"
	flag.StringVar(&cfn, "conf", "./configs/config.yaml", "指定配置文件路径")
	flag.Parse()
	err := config.Init(cfn)
	if err != nil {
		panic(err)
	}

	// init logger
	err = logger.Init(config.Conf.LogConfig, config.Conf.Mode)
	if err != nil {
		panic(err)
	}

	// init Consul
	err = register.Init(config.Conf.ConsulConfig.Addr)
	if err != nil {
		panic(err)
	}

	// listen
	lis, err := net.Listen("tcp", "127.0.0.1:9001")
	//lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", config.Conf.IP, config.Conf.Port))
	if err != nil {
		panic(err)
	}

	// grpc
	s := grpc.NewServer()
	grpc_health_v1.RegisterHealthServer(s, health.NewServer()) // 健康发现
	proto.RegisterSnowFlakeServer(s, &handler.SnowFlakeSrv{})
	go func() {
		err := s.Serve(lis)
		if err != nil {
			panic(err)
		}
	}()

	// 注册服务发现
	err = register.Reg.RegisterServer(config.Conf.Name, config.Conf.IP, config.Conf.Port, nil)
	if err != nil {
		fmt.Println("consul register error", err.Error())
	}
	zap.L().Info("service start")

	// 开启http形式访问
	conn, err := grpc.NewClient(
		fmt.Sprintf("%s:%d", config.Conf.IP, config.Conf.Port),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}
	gwmux := runtime.NewServeMux()
	err = proto.RegisterSnowFlakeHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	gwServer := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.Conf.ApiPort),
		Handler: gwmux,
	}
	zap.L().Info(fmt.Sprintf("Serving gRPC-Gateway on http://0.0.0.0:%d", config.Conf.ApiPort))

	go func() {
		err := gwServer.ListenAndServe()
		if err != nil {
			log.Printf("gwServer.ListenAndServe failed, err: %v", err)
			return
		}
	}()

	// 服务退出时要注销服务
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit // 正常会hang在此处
	// 退出时注销服务
	serviceId := fmt.Sprintf("%s-%s-%d", config.Conf.Name, config.Conf.IP, config.Conf.Port)
	register.Reg.Deregister(serviceId)

}
