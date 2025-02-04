FROM alpine:latest AS builder

RUN apk add --no-cache go=1.23.5-r0

WORKDIR /app
RUN apk add --no-cache git make

COPY . .
RUN go mod download
CMD ["/app/inventoryApp"] 