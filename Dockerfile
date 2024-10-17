FROM golang:1.23

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify
RUN mkdir -p /etc/phoros

COPY ./src ./src
WORKDIR /usr/src/app/src

RUN CGO_ENABLED=0 GOOS=linux go build -o /api
ENV GIN_MODE=release
ENV TZ=UTC

CMD ["/api"]