# Chapter3 <中級編>
## HTTPリクエストの仕方を覚えよう

### 通常問題

1. APIを逐次で実行する

- 案件情報取得APIとユーザ情報取得APIを実行して、
  - ユーザ情報取得APIで取得したユーザIDの案件情報だけをレスポンスとして返してください
  - クライアント ⇄ API ⇄ 案件情報取得API・ユーザ情報取得API
  - 2つのエンドポイントにリクエストしてください（順不同で逐次で実行）
    - ユーザ情報取得API
      - GET
      - http://mock-api/users
      - パラメータは以下を設定してください
        - ヘッダ
          - key: dip
        - クエリパラメータ
          - name: "dip 太郎"
    - 案件情報取得API
      - GET
      - http://mock-api/entries
      - パラメータは以下を設定してください
        - ヘッダ
          - key: dip
      - 下記のレスポンスが返ってくることを確認してください  
      ```
      {
        entries: [
          {
            name: "案件情報1",
            user_id: 123456,
            salary: 123456
          },
          {
            name: "案件情報2",
            user_id: 234567,
            salary: 123456
          }
        ]
      }
      ```


2. 非同期でリクエストする

- 1.で実施したリクエストをgoroutineを使用して、非同期でリクエストしてみましょう
  - 下記のようにすることで非同期で実行することができます
  - 処理結果としては何も返って来なくて問題ありません

``` go
go func() {
  // 処理
}()
```

3. channelを使用する

- 2の改修だけでは何も取得できずに処理が終了してしまいます
  - これはgoroutineの関数が終了する前にmain関数の処理が終了しているためです
  - そういうときはchannelを使いましょう
  - channelの宣言方法は下記のように行います
    - [公式ドキュメント](https://go.dev/doc/effective_go#channels:~:text=we%20need%20channels.-,Channels,-Like%20maps%2C%20channels)

```go

ch := make(chan string, 1)
defer close(ch)

go func() {
  // APIへリクエスト
  res := Request()
  ch <- res
}()

entry := <-ch
// 処理を続ける

```
[参考資料](https://zenn.dev/mikankitten/articles/6344d71f4f4920#channel%E3%81%A8%E3%81%AF)