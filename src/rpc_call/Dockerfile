# 第一阶段：使用官方Golang镜像作为基础镜像来编译应用程序
FROM golang:1.23 AS builder

# 设置工作目录
WORKDIR /app

# 复制go.mod和go.sum文件
COPY go.mod go.sum ./

# 下载依赖
RUN go mod download

# 复制源代码
COPY . .

# 编译Go应用程序
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# 第二阶段：使用Alpine基础镜像来创建一个体积更小的最终镜像
FROM alpine:latest

# 设置工作目录
WORKDIR /app

# 从builder阶段复制编译好的二进制文件
COPY --from=builder /app/main .

# 运行二进制文件
CMD ["./main"]
