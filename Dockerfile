FROM golang:1.24-alpine
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY *.go ./
RUN go build -o /dice-api
EXPOSE 8080
CMD ["/dice-api"]