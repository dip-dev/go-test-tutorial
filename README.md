# go-test-tutorial
Golangで単体テストコードを書くためのチュートリアルです。
既存のコードに対してテストコードを作成してください。

#### ディレクトリ構成
```shell
├── build               : docker buildで利用するファイル。
├── mysql               : DB接続クライアント生成などの処理。当チュートリアルでは単体テストコード作成対象外。
├── tools               : mock生成スクリプト。
└── chapters            : 当チュートリアルで単体テストコードの作成対象となる資産。
    └── chapterX        : チャプター毎のパッケージ。
        ├── chapterX.go : 各チャプターのプログラム。
        ├── main        : chapter4のみ。メイン処理。
        └── data        : chapter4のみ。DBアクセス処理。
```

## 初期構築
- 当リポジトリをForkしてください
- Forkしたリポジトリを任意のディレクトリでclone
- .gitconfigを作成。`$ make create-gitconfig GITHUB_TOKEN=値`
- `$ make up`を実行してコンテナを起動

## チュートリアル
- 課題は`chapters`ディレクトリ配下にchapterで区切られて格納されています。
  - 課題内容は各chapterのREADMEを参照してください。
- テストは以下条件を満たしてください。
  - テーブルドリブンで記述
  - カバレッジ100%到達
- テスト実施手順
  1. `$ make lint`
  1. `$ make gotest`
    - `cover.html`が出力されるので、カバレッジ100%にならない場合はどこのコードが実行されていないのか確認してください。
  - パッケージ単位や関数単位でテストを実行したい場合
    - パッケージ単位：`$ docker compose exec -T <サービス名(※)> go test -v <テスト対象のパッケージまでのパス>`
    - 関数単位：上記パッケージ単位までのコマンドに `-run <関数名>` を追加してください。
      - ※サービス名は`docker-compose.yml`の`services`に記載されているものです。当チュートリアルの場合は`app`になります。

## テストの書き方
```go
// 関数名は必ず「Test」を接頭辞に付ける。
// 引数「t *testing.T」を受けることでテスト関数として動作する。
func TestExample(t *testing.T) {
  // 正常系のテストパターン
  success := map[string]struct {
    paramA string
    paramB int
    want []string
  }{
    "テストケース名": {
      paramA: "abcdefg",
      paramB: 12345,
      want: []string{"a","b","c"},
    },
  }
  // エラー系のテストパターン
  fail := map[string]struct {
    paramA string
    paramB int
    wantErrStr string
  }{
    "テストケース名": {
      paramA: "",
      paramB: -100,
      wantErrStr: "パラメータが不正です",
    },
  }

  for tt, tc := range success {
    t.Run(tt, func(t *testing.T) {
      // テスト対象の関数を呼び出す。
      got, err := example(tc.paramA, tc.paramB)
      // 戻り値が想定通りか検証する。
      if err != nil {
        t.Errorf("test is failed: %s", err)
      }
      if tc.want != got {
        t.Errorf("test is failed: %s", err)
      }
    })
  }
  for tt, tc := range fail {
    t.Run(tt, func(t *testing.T) {
      got, err := example(tc.paramA, tc.paramB)
      if got != nil {
        t.Errorf("test is failed: %s", err)
      }
      if tc.wantErrStr != err.Error() {
        t.Errorf("test is failed: %s", err)
      }
    })
  }
}
```
