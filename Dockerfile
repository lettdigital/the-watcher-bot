FROM golang:alpine3.13 as builder

RUN apk update && apk add libc-dev gcc

WORKDIR /app

COPY . .

RUN go mod tidy

RUN GOARCH=amd64 go build -o app .


FROM alpine:3.13

COPY --from=builder /app/app .

CMD ["./app"]
