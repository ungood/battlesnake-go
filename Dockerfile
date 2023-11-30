# Build stage
FROM golang:alpine as builder

WORKDIR /usr/src

COPY . .
RUN go mod download && go mod verify
RUN go build -v -o /usr/local/bin .

# App stage
FROM alpine:latest
ARG APP=battlesnake-go

RUN addgroup -S app
RUN adduser -S app -G app
COPY --from=builder /usr/local/bin/battlesnake-go /usr/local/bin

USER app
EXPOSE 8080
ENTRYPOINT ["/usr/local/bin/battlesnake-go", "serve"]