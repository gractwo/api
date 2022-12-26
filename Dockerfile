FROM golang:latest
WORKDIR /app
COPY ./ ./
RUN go get .
RUN go run .