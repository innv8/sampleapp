FROM golang:1.22

RUN addgroup users
RUN adduser user
RUN usermod -aG users user

WORKDIR /app
COPY . .
RUN go build -o main main.go
RUN chown user: main
CMD ["./main"]