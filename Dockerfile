FROM golang:1.21.5-alpine
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /example_http
EXPOSE 8080
CMD [ "/example_http" ]