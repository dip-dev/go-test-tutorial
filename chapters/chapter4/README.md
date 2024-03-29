# Chapter4
`Require 100% coverage`

### 実践編2

- 以下プログラムのテストコードを作成してください
- テストコードの作成対象は以下です
  - `/main/main.go`
    - 関数名:`exec`、`validate`
  - `/db/select.go`
    - 関数名:`SelectPrefectures`

## プログラム仕様
- 地方(area)を受け取り、DBを検索した結果を複数件返却する

#### 入力パラメータ

| 論理名 | 物理名 | 型 | 必須 | 最大桁数 |
|:--|:--|:--|:--|:--|
| 地方 | area | 文字列 | ○ | 10 |

#### レスポンス

| 論理名 | 型 | 正常時 | エラー時 |
|:--|:--|:--|:--|
| 検索結果 | Response | ○(別表参照) | nil |
| エラー | error | nil | ◯ |

- 正常時

| No | 論理名 | 物理名 | 型 | 備考 |
|:--|:--|:--|:--|:--|
| 1 | 件数 | Count | 数値 |  |
| 2 | 都道府県リスト | Prefectures | 配列 | 検索結果が0件の場合は空の配列を設定する |
| 2-1 | ID | ID | 数値 |  |
| 2-2 | 名称 | Name | 文字列 |  |
| 2-3 | 地方 | Area | 文字列 |  |
| 2-4 | 面積 | LandArea | 数値 |  |
