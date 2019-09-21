FROM golang:1.12
ARG examplename

WORKDIR /app
COPY . .
RUN ls 
RUN go get
RUN go build -o grpcsample ./$examplename
