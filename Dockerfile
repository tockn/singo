FROM golang:1.14.2-alpine3.11

WORKDIR /app
ADD / /app/
RUN go build -o singo
CMD ./singo
