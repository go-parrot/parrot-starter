# Build stage
FROM golang:1.16.5-alpine3.12 AS builder

# The latest alpine images don't have some tools like (`git` and `bash`).
# Adding git, bash and openssh to the image
RUN apk add --no-cache git make bash ca-certificates tzdata \
    --repository http://mirrors.aliyun.com/alpine/v3.11/community \
    --repository http://mirrors.aliyun.com/alpine/v3.11/main

RUN GRPC_HEALTH_PROBE_VERSION=v0.3.0 && \
    wget -qO/bin/grpc_health_probe \
    https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-linux-amd64 && \
    chmod +x /bin/grpc_health_probe

# 镜像设置必要的环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY="https://goproxy.cn,direct" \
    TZ=Asia/Shanghai \
    APP_ENV=docker

# 移动到工作目录
WORKDIR /go/src/parrot-lagout

# 复制项目中的 go.mod 和 go.sum文件并下载依赖信息
COPY go.mod .
COPY go.sum .

# 将代码复制到容器中
COPY . .
COPY config ./config

# Build the Go app
RUN make build

# Final stage
FROM alpine:3.15

WORKDIR /bin

ENV APP_ENV prod

# 从builder镜像中把 /build 拷贝到当前目录
COPY --from=builder /go/bin/parrot-starter          /bin/parrot-starter
COPY --from=builder /go/src/parrot-starter/config   /data/conf/parrot-starter/config
COPY --from=builder /go/bin/grpc_health_probe 	  /bin/grpc_health_probe

RUN apk update \
    && apk add --no-cache curl jq \
    && rm -rf /var/cache/apk/* \
    && mkdir -p  /data/logs/

# Expose port 8080 and 9090 to the outside world
EXPOSE 8080

# 需要运行的命令
CMD ["/bin/parrot-starter", "-c", "/data/conf/parrot-starter/config"]

# 1. build image: docker build -t qloog/parrot-starter:v1.0.0 -f deploy/docker/Dockerfile .
# 2. start: docker run --rm -it -p 8080:8080 parrot-starter:v1.0.0
# 3. test: curl -i http://localhost:8080/health
# 4. tag: docker tag parrot-starter:v1.0.0 qloog/parrot-starter:v1.0.0
# 5. push tag: docker push qloog/parrot-starter:v1.0.0

