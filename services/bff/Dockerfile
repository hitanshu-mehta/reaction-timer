FROM golang:1.15 as builder
WORKDIR /go/src/github.com/hitanshu-mehta/reaction-timer

COPY services services
COPY go.mod go.mod
COPY go.sum go.sum

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bff services/bff/main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY --from=builder /go/src/github.com/hitanshu-mehta/reaction-timer/bff .

ENTRYPOINT ["/app/bff"]