FROM golang
WORKDIR /app
COPY . .
CMD go run main.go
