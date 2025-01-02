# 環境設定
APP_NAME = app

# タスク定義

# アプリケーションのビルド
build:
	docker-compose build $(APP_NAME)
	docker-compose up -d

# Docker コンテナの起動
up:
	docker-compose up -d

start:
	docker-compose build $(APP_NAME)
	docker-compose up -d
	docker logs -f go_todo_app

# Docker コンテナの停止
down:
	docker-compose down

# アプリケ０ーションのログ表示
logs:
	docker logs -f go_todo_app

# データベースにログイン
db-login:
	docker exec -it go_todo_postgres psql -U user -d go_todo

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
