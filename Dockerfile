FROM golang:1.21-alpine3.18

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify
RUN go mod tidy

COPY . .
RUN go build -v -o /app/bin/toode /app/cmd/main

CMD ["/app/bin/toode"]