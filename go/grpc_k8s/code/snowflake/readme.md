## 安装依赖

```bash
go install \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
    google.golang.org/protobuf/cmd/protoc-gen-go \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc
go get github.com/bwmarrin/snowflake
go get -u google.golang.org/grpc
go get github.com/grpc-ecosystem/grpc-gateway/v2@v2.16.2
```

## 定义proto

```protobuf
syntax = "proto3";

package idSrv;
option go_package = ".;proto";

import "google/protobuf/empty.proto";
import "google/api/annotations.proto";

service SnowFlake {
    rpc GetId (google.protobuf.Empty) returns (Id) {
        option (google.api.http) = {
            get: "/v1/id"
        };
    }
}

message Id {
    int64 id = 1;
}
```

## 生成代码

```bash
protoc \
  -I . -I /home/hsiangya/go/googleapis \
  --go_out . --go_opt paths=source_relative \
  --go-grpc_out . --go-grpc_opt paths=source_relative \
  --grpc-gateway_out . --grpc-gateway_opt logtostderr=true,paths=source_relative \
  proto/snowflake.proto
```

## 安装consul

```bash
# 下载consul
sudo wget https://releases.hashicorp.com/consul/1.18.2/consul_1.18.2_linux_amd64.zip

# 解压后进入目录并赋予权限
chmod 777 consul 

# 拷贝到用户目录
cp consul /usr/local/bin/

# 启动 client后面为服务器ip，启动后默认端口8500，访问8500端口出现ui表示成功
consul agent -dev -ui -node=consul-dev -client=0.0.0.0

```

