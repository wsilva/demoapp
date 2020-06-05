FROM golang:1.13-alpine AS builder
COPY demo.go /src/demo.go
RUN GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -ldflags '-w -s -extldflags "-static"' -o /src/demo /src/demo.go 

FROM scratch
COPY --from=builder /src/demo /demo
EXPOSE 8080
ENTRYPOINT [ "/demo" ]
