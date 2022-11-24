# Builder image
FROM golang:1.18-alpine AS builder

RUN apk add --no-cache gcc musl-dev
COPY . /app
WORKDIR /app
RUN go build -o /app/server .

# Final image
FROM alpine:latest

COPY --from=builder /app/server /app/server
COPY --from=builder /app/blank.db /app/notes.db
WORKDIR /app

EXPOSE 3000
CMD ["/app/server"]
