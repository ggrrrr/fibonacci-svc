DOCKER_REPO ?= "local"
GIT_HASH ?= $(shell git log --format="%h" -n 1)

ab_test:
	ab -t 2s "http://127.0.0.1:8090/v1/next"

dev_env:
	docker-compose up -d

test_svc:
	go test -cover -race ./...

build_svc:
	docker build \
		-f .deploy/Dockerfile \
		--tag "${DOCKER_REPO}/be/fibonacci-svc:${GIT_HASH}" \
		./

tag_svc:
	docker tag "${DOCKER_REPO}/be/fibonacci-svc:${GIT_HASH}" "${DOCKER_REPO}/be/fibonacci-svc:latest"
	
docker_clean:
	docker-compose kill
	docker rm $(docker ps -a -q)

docker_prune:
	docker image prune -a -f
	docker volume prune -a -f