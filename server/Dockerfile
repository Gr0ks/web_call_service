# build golang binary
FROM golang:1.14.0-alpine3.11 AS builder

RUN go version
RUN apk add git

COPY ./ /go/src/backend/
WORKDIR /go/src/backend/

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o ./app ./main.go

# lightweight alpine container with binary
FROM alpine:latest

RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY --from=builder /go/src/backend/cert.pem .
COPY --from=builder /go/src/backend/key.pem .
COPY --from=builder /go/src/backend/app .