FROM golang:1.12 AS builder
WORKDIR /app
COPY . .
RUN go get
RUN go build -o grpcsample ./simple
