FROM golang:1.18 AS builder

COPY ./ /app
WORKDIR /app

RUN echo "" > /app/build/support-files/sql/0001_iam_20200327-1442_mysql.sql
RUN sed -i "1 i -- +migrate Up" /app/build/support-files/sql/*
RUN sed -i 's/`bkiam`.//g' /app/build/support-files/sql/*

ARG BINARY=iam

RUN make build && chmod +x ${BINARY}
RUN mkdir -p /tmp/app/logs
RUN cp ${BINARY} /tmp/app
RUN cp -r /app/build/support-files /tmp/app/support-files

FROM centos:7
COPY --from=builder /tmp/app /app
CMD ["/app/iam", "-c", "/app/config.yaml"]
