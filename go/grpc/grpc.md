# 简介

## 安装

```bash
# 安装依赖
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go get -u google.golang.org/grpc
```

## Protocal Buffers

Protocal Buffers，也就是 protobuf，它是接口设计语言（IDL），它与编程语言无关，可以生成所有主流编程语言的代码，而且，它是二进制格式的数据，比较适合传递大量的数据。

### 语法

```protobuf
syntax = "proto3";

package user;

option go_package = "userpb";
import "enum.proto";

message User {
  int64 id = 1;
  string name = 2;
  repeated string emails = 3;
  Gender gender = 4;
  reserved 3, 16 to 100, 200 to max;
  reserved "uid", "uname";
}
```

![image-20240413182110007](./assets/image-20240413182110007.png)

### 编译器

```bash
protoc --go_out=server --go_opt=paths=source_relative --go-grpc_out=server --go-grpc_opt=paths=source_relative -I proto proto/chat.proto
```

- **--go_out**：指定生成的go代码输出的位置(当前项目往下)

## 启动一个最小服务

```go
package main

import (
	proto "goods/services"
	"google.golang.org/grpc"
	"net"
)

type GoodServer struct {
	proto.UnimplementedGoodsServer
}

func main() {
	// 监听端口
	lis, err := net.Listen("tcp", "0.0.0.0:9001")
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	// 注册服务
	proto.RegisterGoodsServer(s, &GoodServer{})
    
    // 服务与端口dining榜第
	err = s.Serve(lis)
	if err != nil {
		panic(err)
	}
}

```

## 客户端与服务端通信流程

![image-20240716225332792](./assets/image-20240716225332792.png)
