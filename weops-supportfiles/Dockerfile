FROM golang:1.18 AS builder

WORKDIR /codes
COPY . .
RUN VERSION=`head -1 ./VERSION`
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go build -mod=vendor -tags=jsoniter -ldflags "-X iam/pkg/version.Version=${VERSION} -X iam/pkg/version.Commit=`git rev-parse HEAD` -X iam/pkg/version.BuildTime=`date +%Y-%m-%d_%I:%M:%S` -X 'iam/pkg/version.GoVersion=`go version`'" .

FROM centos:7
WORKDIR /app
COPY --from=builder /codes/VERSION ./VERSION
COPY --from=builder /codes/build ./build
COPY --from=builder /codes/release.md ./release.md

RUN mkdir -p /app/bin
COPY --from=builder /codes/iam /app/bin/iam