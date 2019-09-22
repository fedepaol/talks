FROM golang:1.12
ARG examplename

WORKDIR /app
COPY . .
RUN echo $examplename
RUN go get
RUN go build -o grpcsample ./$examplename
