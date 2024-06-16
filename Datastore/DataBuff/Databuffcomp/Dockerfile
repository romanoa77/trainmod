# syntax=docker/dockerfile:1
FROM golang:1.22.4-alpine as build

WORKDIR /app/class
COPY class .

WORKDIR /app/main
COPY main .

#export DOCKER_BUILDKIT=1

RUN --mount=type=cache,target=/root/.cache go mod download

RUN go build -o /app/main/databuff

FROM alpine:latest

ENV ADMROOT=appdata/adm/
ENV ADMFNM=StatDesc.json

ENV DATAROOT=data/

ENV LOGROOT=appdata/log/
ENV LOGSTREAM=LogStream.json


ENV SERVURL=:8080


WORKDIR /app
RUN mkdir appdata
RUN mkdir data
COPY --from=build /app/main/databuff .


CMD  [ "./databuff" ]


