FROM golang:1.17-alpine

RUN apk update
RUN apk add --update alpine-sdk

WORKDIR /app

COPY app/ ./

RUN go mod download
RUN go build -o form3app .

CMD ["go", "test", "-v", "./pkg"]