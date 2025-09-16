# Backend - Go API
タスク管理アプリのバックエンドAPI
## 技術スタック
- Go
- Echo (Webフレームワーク)
- GORM (ORM)
- PostgreSQL
## アーキテクチャ
クリーンアーキテクチャに準拠
```
internal/
├── domain/        # エンティティ・ドメインモデル
├── usecase/       # ビジネスロジック
├── repository/    # データアクセス層
├── handler/       # HTTP ハンドラー（Echo）
└── infrastructure/ # 外部依存（DB接続等）
```
## 開発
```bash
go mod init task-api
go run cmd/main.go
```
http://localhost:8080 でAPI起動