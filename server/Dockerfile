FROM golang:1.9.2 as builder

# Optional WORKDIR (do I need to create the dir?)
RUN mkdir -p /go/src/github.com/nleiva/lab-inventory/
WORKDIR /go/src/github.com/nleiva/lab-inventory/
COPY ./. .

RUN go build -o serv server/main.go

# FROM scratch
FROM gcr.io/distroless/base
COPY --from=builder /go/src/github.com/nleiva/lab-inventory/serv /

CMD ["/serv"]
