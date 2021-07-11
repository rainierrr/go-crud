# go-crud
## 環境構築方法
1. Dockerとdocker-composeのインストール
- [Docker のインストール(公式ドキュメント)](https://docs.docker.jp/engine/installation/index.html)
- [Docker Compose のインストール(公式ドキュメント)](https://docs.docker.jp/compose/install.html)
2. イメージビルド
```
docker-compose build
```
3. パッケージインストール
```
docker-compose run --rm app make install
docker-compose run --rm node yarn install
```
4. migration
```
docker-compose run --rm app sql-migrate up
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

## 開発用の設定
```
git config --local core.hooksPath .githooks
```
settings.json(vscode)
```
{
    "eslint.workingDirectories": [ "./front" ]
}
```
