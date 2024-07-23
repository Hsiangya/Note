# 用户积分系统

## 系统设计

### 功能设计

- 发放积分：单次任务，每日限额
- 扣间积分：兑换扣减 ，惩罚扣减
- 积分过期和衰减：国企清零，周期衰减

### 数据库设计

**积分系统：**

- 任务积分：任务，积分，每日限额，生效时间 
- 用户积分：用户，总积分
- 积分明细：用户，任务，积分，获取时间
- 积分商城：单独设计

**用户等级系统：**

-  等级：等级名称，等级描述，成长数值，有效期
- 特权：等级，产品，功能，有效期
- 用户等级：用户，等级，到期时间，成长数值
- 特价商品：需要时可以额外设计
- 礼品中心：需要时可以额外设计

## 定义pb文件

- 文件目录

```bash
.
├── database
├── dockerfile
└── pb
    └── user_growth.proto
```

- user_growth.proto

```protobuf
syntax = "proto3";
option go_package = "growth/pb";

package UserGrowth;

// 服务
service UserCoin{
  // 获取所有的积分任务列表
  rpc ListTasks(ListTasksRequest) returns (ListTasksReply){}
  // 获取用户的积分信息
  rpc UserCoinInfo(UserCoinInfoRequest) returns (UserCoinInfoReply){}
  // 获取用户的积分明细列表
  rpc UserCoinDetails(UserCoinDetailsRequest) returns (UserCoinDetailsReply){}
  // 调整用户积分-奖励和惩罚都是用这个接口
  rpc UserCoinChange(UserCoinChangeRequest) returns (UserCoinChangeReply){}
}

service UserGrade {
  // 获取所有的等级信息列表
  rpc ListGrades(ListGradesRequest) returns (ListGradesReply){}
  // 获取等级的特权列表
  rpc ListGradePrivileges(ListGradePrivilegesRequest) returns (ListGradePrivilegesReply){}
  // 检查用户是否有某个产品的特权
  rpc CheckUserPrivilege(CheckUserPrivilegeRequest) returns (CheckUserPrivilegeReply){}
  // 获取用户的等级信息
  rpc UserGradeInfo(UserGradeInfoRequest) returns (UserGradeInfoReply){}
  // 调整用户的等级成长值
  rpc UserGradeChange(UserGradeChangeRequest) returns (UserGradeChangeReply){}
}

// 请求和响应消息
message ListTasksRequest{}
message ListTasksReply{
  repeated TbCoinDetail dataLIst = 1;
}
message UserCoinInfoRequest{
  int32 uid = 1;
}
message UserCoinInfoReply{
  TbCoinUser data = 1;
}
message UserCoinDetailsRequest{
  int32 uid = 1;
  int32 page = 2;
  int32 size = 3;
}
message UserCoinDetailsReply{
  repeated TbCoinDetail dataList = 1;
  int32 total = 2;
}
message UserCoinChangeRequest{
  int32 uid = 1;
  string task = 2;
  int32 coin = 3;
}
message UserCoinChangeReply{
  TbCoinUser user = 1;
}

message ListGradesRequest{}
message ListGradesReply{
  repeated TbGradeInfo datalist = 1;
}
message ListGradePrivilegesRequest{
  int32 grade_id = 1;
}
message ListGradePrivilegesReply{
  repeated TbGradePrivilege datalist = 1;
}
message CheckUserPrivilegeRequest{
  int32 uid = 1;
  string product = 2;
  string function = 3;
}
message CheckUserPrivilegeReply{
  bool data = 1;
}
message UserGradeInfoRequest{
  int32 uid = 1;
}
message UserGradeInfoReply{
  TbGradeInfo data = 1;
}
message UserGradeChangeRequest{
  int32 uid = 1;
  int32 score = 2;
}
message UserGradeChangeReply{
  TbGradeUser data = 1;
}


// 数据表模型消息
message TbCoinDetail{
  int32 id = 1;
  int32 uid = 2;
  int32 task_id = 3;
  int32 coin = 4;
  string sys_created = 5;
  string sys_updated = 6;
}
message TbCoinUser{}
message TbGradeInfo{}
message TbGradePrivilege{}
message TbGradeUser{}
```

## 自动生成框架代码、验证服务

1. 自动生成grpc代码
2. 创建目录（`mainserver/mainclient/ugserver`）
3. 实现服务端、客户端main方法
4. 验证服务

- 生成代码

```bash
cd pb
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative user_growth.proto
```

- 创建对应的文件及目录

```bash
├── database
├── dockerfile
├── mainclient
│   └── main.go
├── mainserver
│   └── main.go
├── pb
│   ├── user_growth_grpc.pb.go
│   ├── user_growth.pb.go
│   └── user_growth.proto
└── ugserver
    ├── coin_server.go
    └── grade_server.go
```

- 定义coin_server服务框架

```go
package ugserver

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"growth/pb"
)

type UgCoinServer struct {
	pb.UnimplementedUserCoinServer
}

func (s *UgCoinServer) ListTasks(ctx context.Context, in *pb.ListTasksRequest) (*pb.ListTasksReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "方法待实现")
}
func (s *UgCoinServer) UserCoinInfo(ctx context.Context, in *pb.UserCoinInfoRequest) (*pb.UserCoinInfoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "方法待实现")
}
func (s *UgCoinServer) UserCoinDetails(ctx context.Context, in *pb.UserCoinDetailsRequest) (*pb.UserCoinDetailsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "方法待实现")
}
func (s *UgCoinServer) UserCoinChange(ctx context.Context, in *pb.UserCoinChangeRequest) (*pb.UserCoinChangeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "方法待实现")
}
```

- 定义grade_server服务框架

```go
package ugserver

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"growth/pb"
)

type UgGradeServer struct {
	*pb.UnimplementedUserGradeServer
}

func (s *UgGradeServer) ListGrades(ctx context.Context, in *pb.ListGradesRequest) (*pb.ListGradesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "方法待实现")
}
func (s *UgGradeServer) ListGradePrivileges(ctx context.Context, in *pb.ListGradePrivilegesRequest) (*pb.ListGradePrivilegesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "方法待实现")
}
func (s *UgGradeServer) CheckUserPrivilege(ctx context.Context, in *pb.CheckUserPrivilegeRequest) (*pb.CheckUserPrivilegeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "方法待实现")
}
func (s *UgGradeServer) UserGradeInfo(ctx context.Context, in *pb.UserGradeInfoRequest) (*pb.UserGradeInfoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "方法待实现")
}
func (s *UgGradeServer) UserGradeChange(ctx context.Context, in *pb.UserGradeChangeRequest) (*pb.UserGradeChangeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "方法待实现")
}
```

- 编写mainserver/main.go

```go
package main

import (
	"google.golang.org/grpc"
	"growth/pb"
	"growth/ugserver"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", "127.0.0.1:7789")
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
```

- 编写`mainclient/main.go`

```go
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
```

## 数据层、服务层代码

1. 安装xorm，生成数据模型(`./models`)
2. 服务的数据库配置和连接示例
3. 实现数据层封装(`./dao`)
4. 实现服务层封装（`./server`）

- 安装依赖 cmd：`https://github.com/go-xorm/cmd`

```bash
go get xorm.io/xorm
go get xorm.io/builder
go get xorm.io/reverse
go install xorm.io/reverse@latest
go get -u github.com/go-sql-driver/mysql
```

- 编写数据库连接文件

```yml
---
kind: reverse
name: user_growth
source:
  database: mysql
  conn_str: "hsiangya:9kX=AwM%raN3g?MW@tcp(localhost:31766)/user_growth?charset=utf8"
targets:
  - type: codes
    language: golang
    output_dir: ../models/user_growth/ #  这里依据实际需要生成的目录填写
```

- 执行reverse命令

```bash
reverse -f mysql-usergrowth.yml
```

