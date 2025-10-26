# Start builder
FROM golang:1.24-alpine as builder

WORKDIR /var/www/html

# Copy module files dan download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy seluruh source code
COPY . .

# Default working directory
RUN go mod tidy
RUN go build -o main ./command/api
RUN go build -o worker ./command/worker
# End builder

FROM alpine:3.19

WORKDIR /var/www/html
# Copy binary dari stage builder
COPY --from=builder /var/www/html/main /var/www/html
COPY --from=builder /var/www/html/worker /var/www/html

# Create folder logs
# mkdir all folders
RUN mkdir -p /var/www/html/logs 
RUN touch /var/www/html/logs/activity.log /var/www/html/logs/worker.log

EXPOSE 8080

# Exec command
CMD ["./main"]

USER root