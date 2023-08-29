FROM golang:latest

WORKDIR /app

COPY . .

RUN go build -o myapp

CMD ["./myapp"]
