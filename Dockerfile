FROM golang:1.16-alpine

WORKDIR /app

COPY . .

RUN go build -o dompet

EXPOSE 8080

CMD ./dompet