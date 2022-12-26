FROM docker.io/library/golang:1.18-bullseye

COPY . /go/src/github.com/mingyech/dtls
WORKDIR /go/src/github.com/mingyech/dtls/e2e

CMD ["go", "test", "-tags=openssl", "-v", "."]
