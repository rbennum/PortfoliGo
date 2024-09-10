FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o /portfoligo

FROM alpine:3.19
WORKDIR /app
COPY --from=builder ./portfoligo ./portfoligo
COPY templates templates
COPY dist dist
COPY prod.env prod.env
CMD ["/app/portfoligo"]