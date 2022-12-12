# Chapter3
### 実践編1

- 以下プログラムのテストコードを作成してください

## プログラム仕様
### 関数名：`selectPrefecture`
- 県名(name)を受け取り、DBを検索した結果を1件返す
- 検索結果が存在しない場合は`nil`を返す

#### 入力パラメータ

| 論理名 | 物理名 | 型 |
|:--|:--|:--|
| 県名 | name | 文字列 |

#### レスポンス

| 論理名 | 型 | 正常時(検索結果あり) | 正常時（検索結果なし） | エラー時 |
|:--|:--|:--|:--|:--|
| 検索結果 | *structs.MPrefecture | ○(別表参照) | nil | nil |
| エラー | error | nil | nil | ○ |


- 正常時(検索結果あり) 

| No | 論理名 | 物理名 | 型 |
|:--|:--|:--|:--|
| 1 | ID | ID | 数値 |
| 2 | 名称 | Name | 文字列 |
| 3 | 地方 | Area | 文字列 |
| 4 | 面積 | LandArea | 数値 |


## TestMain(m *testing.M)について
テスト実行時、「m *testing.M」を引数に受ける関数が最初に呼ばれる<br> 
これを利用してテストに必要な前処理・後処理を記載することができる（スコープは各パッケージ単位）<br> 
このTestMain(m *testing.M)自体は必須ではないが、共通的な処理を各テストケースに記載する手間を省くことができる

```go
func TestMain(m *testing.M) {
	// 前処理 start
	err := config.NewConfig()
	if err != nil {
		log.Fatalf("[FATAL] %v\n", err)
	}
	cli, err := mysql.New()
	if err != nil {
		log.Fatalf("[FATAL] %v\n", err)
	}
	mysqlCli = cli
	// 前処理 end

	// テストケース実行
	result := m.Run()

	// 後処理 start
	fmt.Println("[sample]テスト後処理")
	// 後処理 end

	// os.Exitに0が渡ればテスト成功、それ以外の場合は失敗となる。
	os.Exit(result)
}

```

defer（遅延実行）を用いることで後処理は以下の様に書ける
```go

func TestMain(m *testing.M) {
	defer fmt.Println("[sample]テスト後処理")

	err := config.NewConfig()
	if err != nil {
		log.Fatalf("[FATAL] %v\n", err)
	}
	cli, err := mysql.New()
	if err != nil {
		log.Fatalf("[FATAL] %v\n", err)
	}
	mysqlCli = cli
	os.Exit(m.Run())
}
```