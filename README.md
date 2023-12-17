# 動作確認手順

## コンテナの立ち上げ
```
$ docker compose up -d
```

cli, db, graphqlという3つのコンテナが立ち上がります。  
cliコンテナでパンの情報を取得しDBに保存します。次にgraphqlコンテナで保存したデータを取得します。  

## パンの情報を取得・DB保存

cliコンテナにログイン
```
$ docker compose exec cli bash
```

`go run main.go add "エントリーID" "アクセストークン"` という形式でエントリーIDに紐づくパンの情報を取得し、同時にDBに保存します。  
注意. 保存したデータを削除するには、dbコンテナのpostgresSQLにログインし、直接go_appデータベース、breadsテーブルのデータを消してください。
```
$ go run main.go add "entry_id" "access_token"
```

## graphqlサーバーから保存したデータの取得

graphqlコンテナにログインし、サーバーを起動します。
```
$ docker compose exec graphql bash
$ go run server.go
```

ブラウザで`http://localhost:8080/`にアクセスし、GraphQL playgroundを開きます。  
以下のクエリでDBに保存したパンの一覧が取得できます。
```
query findBreads {
   breads {
    	id
      name
      created_at
  }
}
```