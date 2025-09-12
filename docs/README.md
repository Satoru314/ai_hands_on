# Task Management App - プロジェクト設計書

## プロジェクト概要
シンプルなタスク管理アプリケーション

### 機能要件
- タスクの作成、読み込み、更新、削除（CRUD）
- タスクID + タスク内容のみのシンプル構造

## アーキテクチャ

### 全体構造
```
/
├── .github/               # GitHub用設定
├── docs/                  # Claude Code用ドキュメント  
├── app/                   # Next.js フロントエンド
│   ├── src/
│   │   ├── app/          # App Router
│   │   ├── components/   # UIコンポーネント
│   │   └── lib/         # ユーティリティ
│   └── README.md
├── api/                  # Go バックエンド
│   ├── cmd/
│   ├── internal/
│   │   ├── domain/       # エンティティ
│   │   ├── usecase/      # ビジネスロジック
│   │   ├── repository/   # データ層
│   │   ├── handler/      # API層（Echo）
│   │   └── infrastructure/ # DB接続（GORM）
│   └── README.md
└── docker-compose.yml    # PostgreSQL
```

## 技術スタック

### フロントエンド
- **Next.js 15**: App Router使用、src無し構成
- **Tailwind CSS**: スタイリング
- **TypeScript**: 型安全性

### バックエンド  
- **Go**: メイン言語
- **Echo**: Webフレームワーク
- **GORM**: ORM
- **PostgreSQL**: データベース
- **クリーンアーキテクチャ**: 設計パターン

### インフラ
- **Docker Compose**: ローカル開発環境
- **PostgreSQL**: Docker コンテナで実行

## 開発方針
- モノレポ管理
- フロントエンド/バックエンド分離
- シンプルで拡張可能な構造
- ベストプラクティスに準拠