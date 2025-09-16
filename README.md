# Task Management App

シンプルなタスク管理アプリケーション

## 構成

- **app/**: Next.js フロントエンド
- **api/**: Go バックエンドAPI
- **docs/**: プロジェクト設計書

## セットアップ

```bash
# データベース起動
docker-compose up -d

# バックエンド起動
cd api
go run cmd/main.go

# フロントエンド起動
cd app
npm run dev
```

詳細は各フォルダのREADMEを参照してください。