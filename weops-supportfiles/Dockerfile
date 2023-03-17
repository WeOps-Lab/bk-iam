FROM golang:1.18 AS builder

WORKDIR /codes
COPY . .
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN make build

FROM centos:7
WORKDIR /app
COPY --from=builder /codes/VERSION ./VERSION
COPY --from=builder /codes/build ./build
COPY --from=builder /codes/release.md ./release.md

RUN mkdir -p /app/bin
COPY --from=builder /codes/iam /app/bin/iam