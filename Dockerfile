FROM golang:1.18-alpine3.16 as builder

WORKDIR /app

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

COPY . .

RUN go build main.go

FROM alpine

COPY --from=builder /app /app

CMD ["./app/main"]