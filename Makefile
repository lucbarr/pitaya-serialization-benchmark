build-linux:
	@GO111MODULE=on GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o bin/pitaya-serialization-benchmark
	@chmod a+x bin/pitaya-serialization-benchmark

build-linux-proto:
	@GO111MODULE=on GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build --tags=proto -o bin/pitaya-serialization-benchmark
	@chmod a+x bin/pitaya-serialization-benchmark

build-linux-json:
	@GO111MODULE=on GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build --tags=json -o bin/pitaya-serialization-benchmark
	@chmod a+x bin/pitaya-serialization-benchmark

docker-build-proto:
	docker build -t pitaya-serialization-benchmark:proto . --build-arg is_proto="true"

docker-build-json:
	docker build -t pitaya-serialization-benchmark:json . --build-arg is_proto="false"

docker-push-proto: docker-build-proto
	docker tag pitaya-serialization-benchmark:proto lucianobarreira/pitaya-serialization-benchmark:proto
	docker push lucianobarreira/pitaya-serialization-benchmark:proto

docker-push-json: docker-build-json
	docker tag pitaya-serialization-benchmark:json lucianobarreira/pitaya-serialization-benchmark:json
	docker push lucianobarreira/pitaya-serialization-benchmark:json

build-proto:
	go generate ./...
	go build --tags=proto .

build-json:
	go build --tags=json .
