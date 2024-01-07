 #Builder process
FROM golang:1.21.5-alpine as build
RUN mkdir /app
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
RUN go mod tidy
ADD . /app
RUN CGO_ENABLED=0 GOOS=linux go build -o /example_http
#Copies compiled binary and runs the service
FROM alpine:3.19
RUN mkdir /app
COPY --from=build /example_http /app/
COPY --from=build /app/html /app/html

EXPOSE 8080
CMD [ "/app/example_http" ]
