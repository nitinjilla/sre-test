#Dockerfile for getfile, runs on port 3001

FROM golang:1.17-alpine

COPY . /getfile

WORKDIR /getfile

RUN go build -o /getfile

EXPOSE 3001

ENTRYPOINT ["/getfile"]
