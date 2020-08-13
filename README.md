# go-api-practice
go言語 net/httpのAPIサーバ/クライアント練習

# 概要
- `net/http`勉強のために作成したAPIサーバ/クライアント
- トイプログラムなので実用性はありません... :sweat_smile:
    - （例）サーバは永続化していない 

# 準備

```sh:
$ cd client
$ go build
$ cd ../server
$ go build
# APIサーバー起動(localhost:8080)
$ ./server
```

# 使い方
`client/client`からcli経由でAPIを叩く

## create

```sh:
$ ./client create -name Taro -age 20
                                  id         name age
3a353775-8af5-4fad-96dd-24257355d7fc         Taro  20
```

## read

## update

## delete
