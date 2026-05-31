FROM golang:1.20-alpine as builder

WORKDIR /build
RUN   apk add --no-cache upx
COPY  go.mod  .
RUN   go mod download -x
COPY  . .
RUN   go build -trimpath -o bin/mirror ./cmd/mirror
RUN   upx --brute /build/bin/mirror


FROM golang:1.20-alpine as base
WORKDIR    /
COPY       --from=builder /build/bin/mirror mirror
VOLUME     /app
WORKDIR    /app
ENTRYPOINT ["/mirror" ]
