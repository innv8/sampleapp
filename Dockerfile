FROM golang:1.22

RUN addgroup -S nonroot \
    && adduser -S nonroot -G nonroot

USER nonroot

WORKDIR /app
COPY . .
RUN go build -o main main.go
RUN chown user: main
CMD ["./main"]
