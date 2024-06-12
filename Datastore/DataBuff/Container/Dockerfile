FROM golang:1.22.4-alpine as build

WORKDIR /app/class
COPY class .

WORKDIR /app/main
COPY main .


RUN go mod download

RUN go build -o /app/main/databuff

FROM alpine:latest

ENV ADMROOT=adm/
ENV ADMFNM=StatDesc.json

ENV DATAROOT=data/

ENV LOGROOT=log/
ENV LOGSTREAM=LogStream.json


ENV SERVURL=:8080


WORKDIR /app
COPY --from=build /app/main/databuff .

CMD  [ "./databuff" ]


