build-linux:
	@GO111MODULE=on GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o bin/pitaya-serialization-benchmark
	@chmod a+x bin/pitaya-serialization-benchmark

docker-build:
	docker build -t pitaya-serialization-benchmark

docker-tag:
	docker tag game-moderator-proxy:latest ${REPO_URL}:${IMAGE_TAG}
