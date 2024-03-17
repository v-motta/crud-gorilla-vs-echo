FROM golang:1.19

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o restaurant-api

RUN CGO_ENABLED=0 GOOS=linux go build -o /restaurant-api

EXPOSE 9000

# Run
CMD ["/restaurant-api"]