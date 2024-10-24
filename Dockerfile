FROM golang:1.22.8-alpine3.20

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify
RUN go mod tidy

COPY . .

RUN touch /app/.env

RUN go build -v -o /app/bin/toode /app/cmd/main

CMD ["/app/bin/toode"]