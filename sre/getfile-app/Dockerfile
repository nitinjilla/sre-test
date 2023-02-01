#Dockerfile for getfile-app

FROM golang:1.17-alpine AS initial
WORKDIR /getfile-app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go test -o getfile-app
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags '-extldflags "-static"' -o getfile-app

FROM scratch 
WORKDIR /
COPY --from=initial /getfile-app
EXPOSE 3001
ENTRYPOINT["/getfile-app"]
