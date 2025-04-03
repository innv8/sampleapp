FROM golang:1.22

RUN addgroup -S nonroot \
    && adduser -S nonroot -G nonroot

USER nonroot

WORKDIR /app
COPY shape /app/shape
COPY main.go /app/main.go 
COPY go.mod /app/go.mod

RUN go build -o main main.go
RUN chown user: main
CMD ["./main"]
