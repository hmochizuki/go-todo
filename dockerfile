# ベースイメージを使用してビルドフェーズを設定
FROM golang:1.23 AS builder

# 作業ディレクトリを設定
WORKDIR /app

# モジュールファイルをコピーして依存関係をダウンロード
COPY go.mod go.sum ./
RUN go mod download

# ソースコードをコピーしてビルド
COPY . .
RUN go build -o main ./cmd/main.go

# 実行環境として Ubuntu 最新版を使用
FROM ubuntu:22.04

# 必要なツールをインストール
RUN apt-get update && apt-get install -y ca-certificates && apt-get clean

# 作業ディレクトリを設定
WORKDIR /root/

# ビルド済みのバイナリをコピー
COPY --from=builder /app/main .

# 必要なポートを公開
EXPOSE 8080

# アプリケーションを実行
CMD ["./main"]
