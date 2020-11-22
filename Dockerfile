FROM golang:1.14-alpine

RUN apk update
RUN apk add \
  g++ \
  git \
  musl-dev \
  tesseract-ocr-dev

WORKDIR $GOPATH/src/go-ocr
COPY go.mod ./
RUN mkdir -p $GOPATH/src/go-ocr

RUN apk add build-base
COPY . $GOPATH/src/go-ocr

RUN go build -o ocr .
EXPOSE 8080

CMD ["/go/src/go-ocr/ocr"]