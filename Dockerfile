FROM golang:latest AS builder

WORKDIR /src

COPY . .

# Build from the directory containing go.mod
WORKDIR /src/src
RUN CGO_ENABLED=0 go build -o /src/app

FROM scratch
COPY html /html
COPY --from=builder /src/app /app
EXPOSE 8080
ENTRYPOINT ["/app"]