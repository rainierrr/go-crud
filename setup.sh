#!/bin/sh

# イメージビルド
docker-compose build

# パッケージインストール
docker-compose run --rm app make install
docker-compose run --rm node yarn install

#migration
docker-compose run --rm app sql-migrate up

#コンテナ作成
docker-compose up -d
