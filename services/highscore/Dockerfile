FROM golang:1.15 as builder
WORKDIR /go/src/github.com/hitanshu-mehta/reaction-timer/services/highscore

COPY . .

RUN go get -u github.com/golang/dep/cmd/dep
RUN dep init && dep ensure
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .

FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY --from=builder /go/src/github.com/hitanshu-mehta/reaction-timer/services/highscore/highscore .

ENTRYPOINT ["/app/highscore"]