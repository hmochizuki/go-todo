# ベースイメージ
FROM golang:1.23 as builder

# アプリケーションの作業ディレクトリを作成
WORKDIR /app

# モジュールファイルと依存関係をコピー
COPY go.mod go.sum ./
RUN go mod download

# アプリケーションのソースコードをコピー
COPY . .

# アプリケーションをビルド
RUN go build -o main ./cmd/main.go

# 実行環境
FROM debian:bullseye-slim

# 必要なツールをインストール
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

# 作業ディレクトリを設定
WORKDIR /root/

# ビルドされたアプリケーションをコピー
COPY --from=builder /app/main .

# アプリケーションを実行
CMD ["./main"]
