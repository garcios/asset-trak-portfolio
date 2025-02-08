FROM golang:1.22.2-alpine3.19 AS builder

# Set Go env
ENV CGO_ENABLED=0 GOOS=linux
WORKDIR /go/src/app

# Install dependencies
RUN apk --update --no-cache add ca-certificates make protoc

# Build Go binary
COPY Makefile go.mod go.sum ./
RUN go env -w GOPROXY=https://goproxy.io,direct/
RUN make init && go mod download
COPY . .
RUN make tidy

RUN go build -o /go/src/app/service.assetprice  ./cmd/main.go

# Deployment container
FROM scratch

COPY --from=builder /etc/ssl/certs /etc/ssl/certs
COPY --from=builder /go/src/app/service.assetprice /app/service.assetprice
COPY --from=builder /go/src/app/data /app/data

ENTRYPOINT ["/app/service.assetprice"]