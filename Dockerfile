FROM golang:1.13.3 AS builder
WORKDIR /go/src/github.com/ritoon/api-vote/
COPY ./ .
RUN go mod download  
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/github.com/ritoon/api-vote/app .
COPY --from=builder /go/src/github.com/ritoon/api-vote/config.test.yaml .
CMD ["./app"]