#Dockerfile for getfile-app

FROM golang:1.17-alpine AS build
WORKDIR /build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags '-extldflags "-static"' -o getfile-app

FROM scratch 
WORKDIR /
COPY --from=build /build/getfile-app /
EXPOSE 3001
ENTRYPOINT ["/getfile-app"]
