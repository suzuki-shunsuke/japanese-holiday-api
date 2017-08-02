# japanese-holiday-api

日本の休日を取得するAPI。
内閣府が提供するCSVを元にしている。
Goの勉強がてら作ってみた。

* http://www8.cao.go.jp/chosei/shukujitsu/gaiyou.html
* http://www8.cao.go.jp/chosei/shukujitsu/syukujitsu.csv

## Build

```
$ make
```

## API Reference

### Request

```yaml
- method: GET
  url: /holidays
  query:
    from:
      description: 取得開始日(>=)
      type: str
      format: Y-m-d
      required: true
    to:
      description: 取得終了日(<)
      type: str
      format: Y-m-d
      required: true
  response:
  - status_code: 200
    description: 正常終了
    body:
    - date:
        type: str
        format: Y-m-d
      name:
        type: str
      type:
        type: int
  - status_code: 5xx
    description: システムエラー
    body:
      message:
        type: str
        description: エラーメッセージ
  - status_code: 4xx
    description: ユーザーエラー(リクエストのパラメータが不正な場合など)
    body:
      message:
        type: str
        description: エラーメッセージ
```

### Response

number | description
--- | ---
0 | 日曜日
1 | 国民の祝日
2 | 日曜日に当たった国民の祝日の後でその日に最も近い「国民の祝日」でない日(休日扱いとなる)
3 | 国民の祝日ではないが、その前日及び翌日が「国民の祝日」である日(休日扱いとなる)

## License

[MIT](LICENSE)
