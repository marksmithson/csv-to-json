FROM golang:1.13 as builder

WORKDIR /app

COPY . .

RUN ls -al

RUN CGO_ENABLED=0 GOOS=linux go build -a -o main cmd/main.go

# Build the final image

FROM gcr.io/distroless/static

COPY --from=builder /app/main /main

USER 65535:65535

VOLUME /in
VOLUME /out

ENTRYPOINT ["/main"]