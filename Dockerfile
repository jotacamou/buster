FROM golang:latest as builder
LABEL maintainer="JR Camou <jr@camou.org>"
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
ENV GO111MODULE=on
RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -installsuffix cgo -o main .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/main .
COPY --from=builder /app/config .
EXPOSE 8081

CMD ["./main"]
