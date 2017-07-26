# japanese-holiday-api

日本の休日を取得するAPI

* http://www8.cao.go.jp/chosei/shukujitsu/gaiyou.html
* http://www8.cao.go.jp/chosei/shukujitsu/syukujitsu.csv

## 関連API

* Google Calendar API: トークンが必要。
* https://holidays-jp.github.io/
  * 個人がGoogle Calendar APIを使って作ったみたい
  * CSVとJSONに対応
  * クエリで絞込とかは出来ない
  * 自動で定期的にデータを更新
* http://s-proj.com/utils/holiday.html
  * 個人が内閣府のデータに基づいて作っている
* http://ay-sys.com/contents/memo_gcalendar

## API Reference

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
    types:
      type: int[]
      description: 休日の種類
      required: false
      default: [10]
    offset:
      type: int
      description: オフセット
      required: false
      default: 0
    limit:
      type: int
      description: 取得する最大件数
      required: false
      default: 1000
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

### Holiday Type

#### Response

number | description
--- | ---
0 | 日曜日
1 | 国民の祝日
2 | 日曜日に当たった国民の祝日の後でその日に最も近い「国民の祝日」でない日(休日扱いとなる)
3 | 国民の祝日ではないが、その前日及び翌日が「国民の祝日」である日(休日扱いとなる)

#### Request

number | description
--- | ---
10 | 全休日
11 | 日曜を除く全休日

## License

[MIT](LICENSE)
