FROM acr.internal.attains.cn/cloud/go-base-container:v3 as builder

WORKDIR /app

COPY . .

RUN git config -l \
    && git config --global url."git@git.internal.attains.cn:".insteadof "https://git.internal.attains.cn/" \
    && go env -w GOPRIVATE="git.internal.attains.cn" \
    && go env -w GO111MODULE=on \
    && go env -w GOPROXY=https://goproxy.cn,direct \
    && go env \
    && go mod tidy \
    && go build -o server server.go

FROM alpine as runner


WORKDIR /app

COPY --from=builder /app/server ./
COPY --from=builder /app/.env ./

ENV TZ="Asia/Shanghai"
ENV ATTAINS_C=remote

RUN cp -a /etc/apk/repositories /etc/apk/repositories.bak \
    && sed -i "s@https://dl-cdn.alpinelinux.org/@https://repo.huaweicloud.com/@g" /etc/apk/repositories \
    && apk update \
    && apk add alpine-conf \
    && /sbin/setup-timezone -z Asia/Shanghai

EXPOSE 9102

ENTRYPOINT ["/app/server"]