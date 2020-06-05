FROM golang:1.14 AS builder
COPY demo.go /src/demo.go
ARG GOOS=linux
ARG CGO_ENABLED=0
ARG GOARCH=amd64 
RUN go get github.com/prometheus/client_golang/prometheus \
    && go get github.com/prometheus/client_golang/prometheus/promauto \
    && go get github.com/prometheus/client_golang/prometheus/promhttp \
    && go build -ldflags '-w -s -extldflags "-static"' -o /src/demo /src/demo.go 

FROM scratch
COPY --from=builder /src/demo /demo
EXPOSE 8080
ENTRYPOINT [ "/demo" ]
