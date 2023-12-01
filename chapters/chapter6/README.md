# Chapter6
## テストを並列実行してみよう

### 0. はじめに
ℹ️ **テスト関数について**<br>
`t.Run()`で括られている箇所をサブテスト関数と言います<br>
また`TestMain(m *testing.M)`との混同を防ぐため、当チュートリアルでは括られていない箇所を便宜上トップテスト関数と表現します<br>

```golang
func TestSample(t *testing.T) { // トップテスト関数
    cases := map[string]struct {
        // ...

    for tt, tc := range cases {
        t.Run(tt, func(t *testing.T) {  // サブテスト関数 ここから
            // ...
        })                              // サブテスト関数 ここまで
    }
}
```

### 1. まずは直列で実行してみよう

- 以下コマンドでテストを実行してください

```
$ docker compose exec app go test -v -count=1 github.com/dip-dev/go-test-tutorial/chapters/chapter6 -run TestPrintString
```

テストの完了に3秒ほど掛かったはずです

### 2. サブテストを並列実行してみよう
- `TestPrintString`のサブテスト関数に`t.Parallel()`を追加してください
```golang
    for tt, tc := range cases {
        t.Run(tt, func(t *testing.T) {
            t.Parallel() // <-追加
            dispArg(tc.arg)
        })
    }
```

- 再度コマンドからテストを実行してください

以下のようにテストは1秒ほどで完了するようになりましたが、Receivedに同じ文字列が出力されてしまっているかと思います<br>

```
=== RUN   TestPrintString
=== RUN   TestPrintString/print_string_1
=== PAUSE TestPrintString/print_string_1
=== RUN   TestPrintString/print_string_2
=== PAUSE TestPrintString/print_string_2
=== RUN   TestPrintString/print_string_3
=== PAUSE TestPrintString/print_string_3
=== CONT  TestPrintString/print_string_1
=== CONT  TestPrintString/print_string_3
=== CONT  TestPrintString/print_string_2
Received: 何事も成らぬは人の為さぬなり
Received: 何事も成らぬは人の為さぬなり
Received: 何事も成らぬは人の為さぬなり
--- PASS: TestPrintString (0.00s)
    --- PASS: TestPrintString/print_string_2 (1.00s)
    --- PASS: TestPrintString/print_string_1 (1.00s)
    --- PASS: TestPrintString/print_string_3 (1.00s)
PASS
ok  	github.com/dip-dev/go-test-tutorial/chapters/chapter6	1.006s
```

これは`t.Parallel()`の挙動が関わっています<br>

`t.Parallel()`が呼び出されるとそのテストケースは実行を一時停止し、次のテストケースの実行へと移ります<br>
全てのケースが実行されると、`t.Parallel()`によって停止されていたケースが並列に再開します<br>

つまり上記ログを確認すると、
- `=== RUN   TestPrintString/print_string_1`で`print_string_1`のケースを開始
- `=== PAUSE TestPrintString/print_string_1`で`print_string_1`のケースを一時停止
- `=== RUN   TestPrintString/print_string_2`で`print_string_2`のケースを開始
- `=== PAUSE TestPrintString/print_string_2`で`print_string_2`のケースを一時停止
- `=== RUN   TestPrintString/print_string_3`で`print_string_3`のケースを開始
- `=== PAUSE TestPrintString/print_string_3`で`print_string_3`のケースを一時停止
- `=== CONT  xxx`のログで一斉にケースを再開
という挙動をしていることが分かります。

それではなぜReceivedに同じ文字列が出力されてしまったのかについてですが、<br>
`for...range`で取り出した変数はループ毎にポインタ(※)が変わらず同じ場所を参照し続けます。<br>
そのためforループ終了後にケースが一時停止されることによって最終ループの値が変数に残り、その値が全てのケースで参照されてしまっていたから、ということになります<br>
（※ポインタについては https://go-tour-jp.appspot.com/moretypes/1 を参照ください）<br>

これを回避するには以下のように**変数のシャドーイング**を行います<br>

- `TestPrintString`の関数実行部を以下のように修正して再度コマンドからテストを実行してください
```golang
    for tt, tc := range cases {
        tc := tc // <-追加。新たに変数tcを定義することで別のポインタに値が格納されるため、値が上書きされなくなる
        t.Run(tt, func(t *testing.T) {
            t.Parallel()
            dispArg(tc.arg)
        })
    }
```

ケース毎の文字列が出力されることが確認できたらOKです<br>

### 3. トップテスト関数を並列実行してみよう
- 以下コマンドでテストを実行してください
```
$ docker compose exec app go test -v -count=1 github.com/dip-dev/go-test-tutorial/chapters/chapter6
```

実行に2秒ほど掛かったかと思います<br>
これはトップテスト関数が直列で実行されているためです<br>

- `TestPrintString`と`TestPrintInt`のトップテスト関数に`t.Parallel()`を追加してください

こうすることで、コード上ではパッケージ内の`t.Parallel()`を呼び出しているすべてのサブテスト関数が、一斉に並列に動作するようになりました。<br>

ただし、これだけでは不十分で実際には**同時に動作するテストの個数は制限**されています。<br>
どれだけの個数のテスト関数が並列に動作するかは、`-parallel`フラグで指定します。<br>
（詳細は https://pkg.go.dev/cmd/go/chapters/test を参照）

- 以下コマンドでテストを実行してください
  - トップテスト関数が2つでそれぞれサブテストが3つ起動するので`-parallel=6`を指定しています
```
$ docker compose exec app go test -v -parallel=6 -count=1 github.com/dip-dev/go-test-tutorial/chapters/chapter6
```

2つのトップテスト関数の実行が1秒ほどで完了したかと思います

並列実行することでテストの実行時間を短縮することができました<br>

### まとめ
ℹ️ **並列性を最大限に向上させるには**<br>
- トップテスト関数とその中のサブテスト関数の両方で、`t.Parallel()`メソッドを呼び出す
- `-parallel`フラグで並列レベルを適切に設定する

⚠️ **並列実行における注意事項**<br>
- `for range`で複数ケースのテストを実行する場合は、変数のシャドーイングを行う
- `-parallel`フラグについて
  - 一度に並列処理できる数は実行環境のCPUに依存するため、<br>`-parallel`の値を大きくすれば必ずしもテスト性能が向上する、というわけではありません
