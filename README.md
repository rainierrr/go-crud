# go-crud
## 環境構築方法
1. Dockerとdocker-composeのインストール

2. イメージビルド
```
docker-compose build
```
3. パッケージインストール
```
docker-compose run --rm app make install
```

4. コンテナ作成
```
docker-compose up -d
```
5. コンテナの起動確認
```
docker-compose ps
```

6. サイトにアクセス
localhost:8080にアクセスしてください。