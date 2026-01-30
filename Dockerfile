FROM golang:1.24.11-alpine AS build

ARG VERSION

ENV GOPROXY=https://goproxy.cn,direct

WORKDIR /root

COPY . /root

RUN sed -i "s/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g" /etc/apk/repositories \
    && apk upgrade && apk add --no-cache --virtual .build-deps \
    ca-certificates upx tzdata

RUN CGO_ENABLED=0 go build --ldflags="-X main.Version=${VERSION}" -o ai-model cmd/ai-model/main.go \
    && chmod +x ai-model

FROM alpine:3.19

COPY --from=build /root/ai-model /app/ai-model

RUN sed -i "s/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g" /etc/apk/repositories \
    && apk upgrade && apk add --no-cache --virtual .build-deps \
    ca-certificates upx tzdata

WORKDIR /app

ENTRYPOINT ["/app/ai-model"]
