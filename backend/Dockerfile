FROM golang:1.18

WORKDIR /app

COPY . .

RUN go build -o openpager

EXPOSE 8080

CMD ["./openpager"]
