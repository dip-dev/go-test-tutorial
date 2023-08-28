# Chapter7
## ヘルパー関数と後処理

### 1. まずはテストを実行してみよう
- テスト関数`TestCreateTextFile()`が呼び出している`deleteTextFile()`に`t.Fatal()`を追加して、テストが失敗するようにしてください
```golang
func deleteTextFile(t *testing.T, target string) {
    t.Fatal("test failed")  // ★追加
    dir, err := os.Getwd()
    // ...
```
- 以下コマンドでテストを実行してください
```
$ docker-compose exec app go test -count=1 github.com/dip-dev/go-test-tutorial/chapters/chapter7
```

以下のような結果になったかと思います
```
--- FAIL: TestCreateTextFile (1.00s)
    --- FAIL: TestCreateTextFile/成功 (0.00s)
        chapter7_test.go:12: test failed
FAIL
FAIL	github.com/dip-dev/go-test-tutorial/chapters/chapter7	1.008s
FAIL
```

### 2. ヘルパー関数としてマークしてみよう
- `deleteTextFile()`関数に`t.Helper()`を追加してヘルパー関数としてマークしてください
```golang
func deleteTextFile(t *testing.T, targets ...string) {
    t.Helper() // ★先頭に追加
    t.Fatal("test failed")
    dir, err := os.Getwd()
    // ...
```
- 再度コマンドからテストを実行してください

以下のような結果になったかと思います
```
--- FAIL: TestCreateTextFile (1.01s)
    --- FAIL: TestCreateTextFile/成功 (0.00s)
        chapter7_test.go:36: test failed
FAIL
FAIL	github.com/dip-dev/go-test-tutorial/chapters/chapter7	1.018s
FAIL
```

エラー箇所の行数が変わっていることが確認できます<br>

> ℹ️ **Note**<br>
> `t.Helper()`を呼び出してヘルパー関数としてマークすることで、<br>
> 表示されるエラー箇所が呼び出し元のテスト関数の行になるため、失敗したテストが特定しやすくなります<br>
> また、エラーが発生し得ない場合でも`t *testing.T`を引数に受けて`t.Helper()`を記述することでテスト用の関数であることが明示的に分かるメリットがあります<br>

### 3. テストの後処理を書き換えてみよう
現状の実装では`createTextFile()`がエラーを返すと後処理である`deleteTextFile()`が実行されません。<br>
ここで確実に実行されるように遅延実行`defer`を使ってみましょう<br>

#### 3-1. deferを使った後処理の実行
- 最初に`deleteTextFile()`に追加した`t.Fatal()`を削除してください
- テスト関数`TestCreateTextFile()`を以下のように書き換えてください
```golang
func TestCreateTextFile(t *testing.T) {

    fileNames := []string{"file1", "file2", "file3"}

    t.Run("成功", func(t *testing.T) {
        for i, fileName := range fileNames {
            defer deleteTextFile(t, fileName) // ★追加
            t.Run(fmt.Sprintf("create file%d", i), func(t *testing.T) {
                err := createTextFile(fileName)
                assert.NoError(t, err)
            })
            // deleteTextFile(t, fileName) ★削除
        }
    })
}
```

- 再度コマンドからテストを実行してください

テストが正常に実行され、出力ファイルが削除されているかと思います<br>

ここで、テストの実行に3秒ほど掛かってしまっているのでchapter6で触れた並列化を実施してみましょう<br>

#### 3-2. 並列化
- `t.Parallel()`を使ってテストを並列化し、再度コマンドからテストを実行してください
  - 並列化の方法についてはchapter6を参照し、自力で実装してみてください

以下のようにテストが失敗したはずです<br>
```
--- FAIL: TestCreateTextFile (1.02s)
    --- FAIL: TestCreateTextFile/成功 (0.00s)
        chapter7_test.go:38: remove /go/src/github.com/dip-dev/go-test-tutorial/chapters/chapter7/file3.txt: no such file or directory
        panic.go:540: remove /go/src/github.com/dip-dev/go-test-tutorial/chapters/chapter7/file2.txt: no such file or directory
        panic.go:540: remove /go/src/github.com/dip-dev/go-test-tutorial/chapters/chapter7/file1.txt: no such file or directory
FAIL
FAIL	github.com/dip-dev/go-test-tutorial/chapters/chapter7	1.025s
FAIL
```

chapter6の以下文言を思い出してください<br>
> `t.Parallel()`が呼び出されるとそのテストケースは実行が一時停止され、次のテストケースの実行へと移ります

テストケースが一時停止されたことで、ファイルが作成される前に`defer`に指定した`deleteTextFile()`が実行されてしまったため上記エラーが発生しています<br>

#### 3-3. t.Creanup()を使う
- `defer`で呼び出している`deleteTextFile()`を以下のように書き換えてください
```golang
// defer deleteTextFile(t, fileName) ★削除
t.Cleanup(func() { 
    deleteTextFile(t, fileName)
})
```

- 再度コマンドからテストを実行してください

エラーが解消され、テストの実行時間も短縮できたかと思います。<br>

> ℹ️ **Note**<br>
> `t.Creanup()`はGo1.14から追加されたメソッドで、<br>
> `t.Cleanup()`メソッドで登録した関数は、すべてのサブテストが終了した時点で呼び出されます<br>
> つまり、`t.Parallel()`でサブテストを並列化している場合、後処理には`t.Cleanup()`を用いた方が無難(※)です<br>
> 
> ※必ずしも`t.Cleanup()`を使わなければならないわけではない。例えばサブテスト関数が、さらに`t.Run()`でネストした「サブテストのサブテスト関数」を含んでいなければ、`defer`でもOK

### まとめ
- ヘルパー関数には`t.Helper()`を使ってマークする
- `t.Parallel()`でサブテストを並列化している場合、後処理には`t.Cleanup()`を用いると良い
