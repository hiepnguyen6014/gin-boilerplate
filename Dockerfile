# Stage 1: compile the program
FROM golang:1.19.4 as build-stage
WORKDIR /app
COPY go.* .
RUN go mod download
RUN go mod tidy

COPY . .

COPY .env .env

RUN go build -o server.exe main.go

# Stage 2: build the image
FROM alpine:latest  
RUN apk --no-cache add ca-certificates libc6-compat
WORKDIR /app/
COPY --from=build-stage /app/server.exe .
COPY --from=build-stage /app/.env .

EXPOSE 8080

CMD ["./server.exe"]  