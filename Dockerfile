FROM golang:1.14.2-alpine3.11

WORKDIR /app
COPY main.go /app
COPY server.go /app
COPY repository /app/repository
COPY model /app/model
COPY handler /app/handler
COPY manager /app/manager
COPY go.mod /app
COPY go.sum /app

WORKDIR /app/demo/dist
COPY demo/dist /app/demo/dist

WORKDIR /app
RUN ls
RUN go build -o singo
ENTRYPOINT ["./singo"]
CMD ["-demo=false"]
