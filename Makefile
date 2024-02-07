DOCKER_REPO ?= local
GIT_HASH ?= $(shell git log --format="%h" -n 1)

ab_test_pg: docker_run_dev_svc
	ab -c100 -t 2s "http://127.0.0.1:8091/v1/next"

ab_test_redis: docker_run_dev_svc
	ab -c100 -t 2s "http://127.0.0.1:8092/v1/next"

ab_test_ram:
	ab -c100 -t 2s "http://127.0.0.1:8093/v1/next"

ab_test_local:
	ab -c100 -t 2s "http://127.0.0.1:8090/v1/next"

dev_env:
	docker-compose up -d

test_svc: docker_run_dev_env
	go test -cover -race ./...

build_svc: test_svc
	docker build \
		-f .deploy/Dockerfile \
		--tag "${DOCKER_REPO}/be/fibonacci-svc:${GIT_HASH}" \
		.
	docker tag "${DOCKER_REPO}/be/fibonacci-svc:${GIT_HASH}" "${DOCKER_REPO}/be/fibonacci-svc:latest"

	
docker_run_dev_env:
	docker-compose up -d redis postgres 

docker_run_dev_svc: docker_run_dev_env
	docker-compose up -d fi_redis fi_pg fi_ram

docker_clean:
	docker-compose kill
	docker rm `docker ps -aq`

docker_prune:
	docker image prune -a -f
	docker volume prune -a -f
