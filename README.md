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

```sh:
$ ./client read
                                  id         name age
fbf8f037-07c5-4ac5-bcf2-5a8d8c860b41         Taro  20
eda12a56-c7d6-4309-b65e-b95525ae1db8         John  25
```

## update

```sh:
$ ./client update -id 419b9dd6-d54a-434c-9239-0144177c117f -name jiro -age 21
                                  id         name age
419b9dd6-d54a-434c-9239-0144177c117f         jiro  21
```

## delete
