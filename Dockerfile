FROM golang:alpine AS build-env
ARG is_proto
RUN echo $is_proto

ADD . /src
RUN apk add make
RUN if [[ "$is_proto" == "true" ]]; then echo "building with protos"; cd /src && make build-linux-proto; else echo "building with json"; cd /src && make build-linux-json; fi

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

FROM alpine

WORKDIR /app
COPY --from=build-env /src/bin/pitaya-serialization-benchmark /app
ADD ./examples /app/examples

CMD /app/pitaya-serialization-benchmark
