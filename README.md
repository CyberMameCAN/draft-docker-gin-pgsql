# Docker + Gin + PostgreSQL + pgAdmin4

Dockerを使ってGinとPostgreSQLなApp開発のベース環境を作っています。

## 参考
- [go-todo](https://github.com/anshul-repos/go-todo)
- [Golang + PostgreSQLの開発環境をDockerで構築](https://koredana.info/blog/docker-golang-postgresql/)

## 使い方

### 起動
    $ docker-compose up --build

### 停止
    $ docker-compose down -v

## Web App

ブラウザでアクセスする。初期投入データをJSONで返してきます。

    http://localhost:7070/api/v1/tasks/
    
## データベース管理, pgAdmin4

ポート番号は**35433**で初期設定しています。  
ブラウザでアクセスしてください。

    http://localhost:35433/