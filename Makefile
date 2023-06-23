.PHONY: build

APP=app

EXEC_APP=docker-compose exec $(APP)
GO_VER=1.20

up:
	docker-compose up -d

stop:
	docker-compose stop

down:
	docker-compose down

build:
	$(EXEC_APP) go mod tidy -go=$(GO_VER)

lint:
	@make build
# デフォルトのタイムアウト値が1mのため、チェックするファイル数が多くなった時のために3mで設定
	$(EXEC_APP) golangci-lint run --fix --timeout 3m

gotest:
	@make lint
	$(EXEC_APP) go test -coverprofile=cover.out ./...
	$(EXEC_APP) go tool cover -html=cover.out -o cover.html

reload:
	docker-compose stop $(APP)
	@make up