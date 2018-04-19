FROM golang:1.9.2 as builder

# Optional WORKDIR (do I need to create the dir?)
RUN mkdir -p /go/src/github.com/nleiva/lab-inventory/
WORKDIR /go/src/github.com/nleiva/lab-inventory/
COPY ./. .

# Disable crosscompiling 
ENV CGO_ENABLED=0

# Ccompile Linux only
ENV GOOS=linux

# Build the binary with debug information removed
RUN go build -ldflags '-w -s' -a -installsuffix cgo -o web webserver/main.go

FROM alpine
COPY --from=builder /go/src/github.com/nleiva/lab-inventory/web /

EXPOSE 8081

CMD ["/web"]