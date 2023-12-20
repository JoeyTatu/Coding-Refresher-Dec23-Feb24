FROM golang:1.12.0-alpine3.9

# Install Git
RUN apk update && apk add git

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN go build -o main .
CMD ["/app/main"]
