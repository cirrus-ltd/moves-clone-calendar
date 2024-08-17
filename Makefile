# Makefile

# Goのパラメータ
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=myapp
BINARY_UNIX=$(BINARY_NAME)_unix

# マイグレーションのパラメータ
MIGRATE_CMD=go run cmd/cli/main.go

# Docker Composeのパラメータ
DOCKER_COMPOSE=docker compose

# 引数なしでmakeが実行されたときのデフォルトターゲット
default: build

# プロジェクトのビルド
build:
	$(GOBUILD) -o $(BINARY_NAME) -v ./cmd/main

# プロジェクトのクリーン
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)

# テストの実行
test:
	$(GOTEST) -v ./...

# マイグレーションの実行
migrate:
	$(MIGRATE_CMD)

# アプリケーションの実行
run:
	./$(BINARY_NAME)

# 依存関係のインストール
deps:
	$(GOGET) -u ./...

# Linux用にクロスコンパイル
build-linux:
	GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v ./cmd/main

# Docker Composeコマンド
docker-up:
	$(DOCKER_COMPOSE) up -d

docker-down:
	$(DOCKER_COMPOSE) down

docker-build:
	$(DOCKER_COMPOSE) build

docker-restart:
	$(DOCKER_COMPOSE) restart

docker-logs:
	$(DOCKER_COMPOSE) logs -f

docker-ps:
	$(DOCKER_COMPOSE) ps
