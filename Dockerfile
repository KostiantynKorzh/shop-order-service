FROM golang:1.16-alpine

WORKDIR /build

COPY src/go.mod .
COPY src/go.sum .
RUN go mod download

COPY src .

RUN go build -o main .

WORKDIR /dist

RUN cp /build/main .

EXPOSE 1313

CMD ["/dist/main"]