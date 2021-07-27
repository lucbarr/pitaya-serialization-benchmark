FROM golang:alpine AS build-env

ADD . /src
RUN apk add make
RUN cd /src && make build-linux

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

FROM alpine

WORKDIR /app
COPY --from=build-env /src/bin/pitaya-serialization-benchmark /app

CMD /app/pitaya-serialization-benchmark
