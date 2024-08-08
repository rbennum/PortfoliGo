FROM golang:1.22-alpine

WORKDIR /app

COPY go.mod ./

# If go.sum actually exists, don't forget to add below code as well
# COPY go.sum ./

RUN go mod download

COPY . .

RUN go build -o /portfoligo

EXPOSE 8080

CMD ["/portfoligo"]