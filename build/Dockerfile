FROM golang:latest as builder
RUN mkdir /resources
ADD . /resources/
WORKDIR /resources

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

RUN go mod download
RUN go build -a -installsuffix cgo -o app cmd/main.go
RUN ls

FROM alpine:latest as cacerts
RUN apk --no-cache add ca-certificates

FROM scratch
EXPOSE 8080
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=cacerts /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /resources/app .
CMD ["./app"]