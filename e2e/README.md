# E2E Testing

Task Management アプリケーションのEnd-to-Endテストディレクトリです。

## 📁 構成

```
e2e/
├── scripts/test_api.sh                        # 自動APIテスト
├── postman/                                   # Postman用ファイル
│   ├── Task_Management_API.postman_collection.json
│   └── Task_Management_API.postman_environment.json
└── README.md                                  # このファイル
```

## 🚀 テスト実行方法

### 1. 事前準備
```bash
# PostgreSQL起動
docker-compose up -d

# API サーバー起動
cd api && go run cmd/main.go
```

### 2. 自動テスト実行（推奨）
```bash
# 実行権限付与（初回のみ）
chmod +x e2e/scripts/test_api.sh

# E2Eテスト実行
./e2e/scripts/test_api.sh
```

### 3. 手動cURLテスト

#### タスク作成
```bash
curl -X POST http://localhost:8080/api/v1/tasks \
  -H "Content-Type: application/json" \
  -d '{"content": "買い物に行く"}'
```

#### 全タスク取得
```bash
curl http://localhost:8080/api/v1/tasks
```

#### 特定タスク取得
```bash
curl http://localhost:8080/api/v1/tasks/1
```

#### タスク更新
```bash
curl -X PUT http://localhost:8080/api/v1/tasks/1 \
  -H "Content-Type: application/json" \
  -d '{"content": "買い物完了", "status": "completed", "due_date": null}'
```

#### タスク削除
```bash
curl -X DELETE http://localhost:8080/api/v1/tasks/1
```

#### ヘルスチェック
```bash
curl http://localhost:8080/health
```

### 4. Postmanテスト

1. Postmanを開く
2. "Import" ボタンをクリック  
3. `e2e/postman/Task_Management_API.postman_collection.json` を選択
4. `e2e/postman/Task_Management_API.postman_environment.json` を選択
5. コレクションを実行

## 📋 テスト対象エンドポイント

- ✅ `GET /health` - ヘルスチェック
- ✅ `POST /api/v1/tasks` - タスク作成
- ✅ `GET /api/v1/tasks` - 全タスク取得
- ✅ `GET /api/v1/tasks/:id` - 特定タスク取得
- ✅ `PUT /api/v1/tasks/:id` - タスク更新
- ✅ `DELETE /api/v1/tasks/:id` - タスク削除

## 📝 ステータス値

- `pending` - 新規作成時のデフォルト
- `in_progress` - 作業中
- `completed` - 完了
- `cancelled` - キャンセル

## 💡 トラブルシューティング

### よくあるエラー
- **Connection refused**: APIサーバーが起動していない
- **Database error**: PostgreSQLが起動していない  
- **jq command not found**: `sudo apt install jq` でインストール

### デバッグ用
```bash
# API サーバーログ確認（APIサーバー起動時のターミナルを確認）

# データベース接続確認
curl http://localhost:8080/health
```