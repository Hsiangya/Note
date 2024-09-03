# 使用Go语言的Alpine镜像
FROM golang:1.22.4-alpine3.20 as builder

# 设置工作目录
WORKDIR /app

# 安装构建依赖
RUN apk add --no-cache git

# 复制go mod和sum文件
COPY go.mod ./
COPY go.sum ./

# 下载所有依赖
RUN go mod download

# 复制源代码到容器中
COPY . .

# 构建应用程序
RUN CGO_ENABLED=0 GOOS=linux go build -v -o server

# 使用Alpine作为运行时环境
FROM alpine:3.20

# 拷贝构建好的程序到新容器
COPY --from=builder /app/server /server

# 运行gRPC服务器
ENTRYPOINT ["/server"]