# 環境設定
APP_NAME = myapp

# タスク定義

# アプリケーションのビルド
build:
	go build -o bin/$(APP_NAME) ./cmd/main.go

# アプリケーションの実行
run:
	go run ./cmd/main.go

# Docker コンテナの起動
up:
	docker-compose up -d

# Docker コンテナの停止
down:
	docker-compose down

# テストの実行
test:
	go test ./...

# 静的解析（例: lint チェック）
lint:
	golangci-lint run

# クリーンアップ
clean:
	rm -rf bin/
	docker-compose down -v
