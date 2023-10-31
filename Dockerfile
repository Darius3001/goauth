FROM golang:1.18

WORKDIR /app

COPY . .

RUN go build -o goauth

EXPOSE 8080

CMD ["./goauth"]
