FROM golang:1-alpine as builder

WORKDIR /app

COPY . ./
RUN go mod download


RUN go build -v -o data-counter

FROM scratch

COPY --from=builder /app/data-counter /app/data-counter

CMD ["/app/data-counter"]
