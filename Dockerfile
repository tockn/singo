FROM golang:1.14.2-alpine3.11

WORKDIR /app
COPY main.go /app
COPY repository /app/repository
COPY model /app/model
COPY handler /app/handler
COPY manager /app/manager
COPY go.mod /app
COPY go.sum /app

WORKDIR /app/example/dist
COPY example/dist /app/example/dist

WORKDIR /app
RUN ls
RUN go build -o singo
CMD ./singo
