ARG GO_VERSION=1.13.7

FROM golang:$GO_VERSION-alpine AS build

RUN apk add --no-cache \
    gcc \
    git \
    musl-dev \
    nodejs \
    openssl \
    postgresql-client \
    yarn

WORKDIR /flipt

COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download

COPY . .
RUN cd ./ui && yarn install && yarn run build
RUN go get -v github.com/gobuffalo/packr/packr
RUN packr -v -i cmd/flipt
RUN go install -v ./cmd/flipt/

RUN echo "cp /go/bin/flipt /out/flipt" > out

CMD [ "/bin/sh", "/flipt/out" ]
