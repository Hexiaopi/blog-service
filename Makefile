.PHONY: base build run test clean swagger init ui

GIT_VERSION = $(shell git describe --tags)
GIT_COMMIT = $(shell git rev-parse --short HEAD)
GIT_BRANCH = $(shell git rev-parse --abbrev-ref HEAD)
GIT_STATE = $(shell [ -z $(git status --porcelain) ] && echo "dirty" || echo "clean")
BUILD_DATE = $(shell date -u +'%Y-%m-%dT%H:%M:%SZ')

LDFLAGS = -X github.com/hexiaopi/blog-service/cmd/server/version.gitVersion=$(GIT_VERSION) \
	-X github.com/hexiaopi/blog-service/cmd/server/version.gitCommit=$(GIT_COMMIT) \
	-X github.com/hexiaopi/blog-service/cmd/server/version.gitBranch=$(GIT_BRANCH) \
	-X github.com/hexiaopi/blog-service/cmd/server/version.gitTreeState=$(GIT_STATE) \
	-X github.com/hexiaopi/blog-service/cmd/server/version.buildDate=$(BUILD_DATE)

PKGS = $(shell go list ./...)
GOFILES = $(shell find . -name "*.go" -type f -not -path "./vendor/*")

base: clean test swagger fmt build

version: 
	@echo '"$(GIT_VERSION)"

help: ## display the help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

build: ## build server
	go build -ldflags "$(LDFLAGS)" -o bin/blog cmd/*.go

generate: ## generate server
	go generate ./...

run: ## run server
	go run cmd/root.go

test: ## run unit test
	go test -ldflags -s -v -coverprofile=cover.out $(PKGS)
	go tool cover -func=cover.out -o coverage.txt

clean: ## clean bin and go mod
	@rm -rf bin/
	go mod tidy

fmt: ## golang format
	gofmt -s -w $(GOFILES)

swagger: ## swagger init
	swag init -g cmd/root.go -o internal/routers/api/docs

init: install-swagger postgres redis ## install all dependencies and init config
	git config core.hooksPath .githooks
	echo "init all"

install-swagger: ## install swagger from golang
	go install github.com/swaggo/swag/cmd/swag@latest

postgres: ## init postgres db
	@docker start mypostgres || docker run --name mypostgres -d -p 5432:5432 -e POSTGRES_PASSWORD=123456 postgres
	until docker exec mypostgres psql -U postgres; do echo "wait postgres start"; sleep 1; done
	cat scripts/db.sql | docker exec -i mypostgres psql -U postgres

exec-db: ## exec to db container
	docker exec -it mypostgres psql -d weave -U postgres

redis: ## init redis
	@docker start myredis || docker run --name myredis -d -p 6379:6379 redis --appendonly yes --requirepass 123456

ui: ## run ui locally
	cd web && npm i && npm run dev

SERVER_IMG=blog-server
docker-build-server: ## build server image
	docker build -t $(SERVER_IMG) .

docker-run-server: ## run server in docker
	docker run --network host -v $(shell pwd)/configs:/configs -v /var/run/docker.sock:/var/run/docker.sock $(SERVER_IMG)

FRONENT_IMG=blog-fronent
docker-build-ui: ## build frontend image
	cd web && docker build -t $(FRONENT_IMG) .

docker-run-ui: ## run frontendx in docker
	docker run --network host -v $(shell pwd)/web/nginx.conf:/etc/nginx/conf.d/nginx.conf $(FRONENT_IMG)
