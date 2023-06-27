FROM golang:alpine AS builder
WORKDIR /workspace
COPY ../. .
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build

FROM alpine:latest
WORKDIR /opt/boundary-census
COPY --from=builder /workspace/boundary-census .
ENTRYPOINT ["./boundary-census"]