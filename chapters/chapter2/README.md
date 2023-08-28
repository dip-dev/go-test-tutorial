# Chapter2
`Require 100% coverage`

### `gomock`の使い方を覚えよう
- 以下条件を満たすように、関数`exec`のテストコードを完成させてください
  - `/chapter2/communication`パッケージのメソッド`Greeting`をmock化し、`Nice to meet you!!`を返却させる

※ `mock`とは  
簡単に言うとメソッドの動作をシミュレートするための仕組みです<br> 
テスト対象のパッケージが呼び出しているメソッドをMockに差し替えることで、<br> 
メソッドの動作を任意に指定したり、メソッドの呼び出し有無や想定した引数が渡されているかを検証することができます。<br> 

> ℹ️ **Note**<br>
> 公式のgomockが2023/6/28にアーカイブされたため、後継となる以下リポジトリのパッケージを参照してください。<br> 
> https://pkg.go.dev/go.uber.org/mock/gomock<br>

## プログラム仕様
### 関数名:`exec`
- communicationパッケージの関数`Greeting`を呼び出し、その戻り値を返す

#### 入力パラメータ
- なし

#### レスポンス
| 論理名 | 型 | 正常時 |
|:--|:--|:--|
| 実行結果 | string | ○ |


## mockの書き方
```go
  // mock contoroller生成
  ctrl := gomock.NewController(t)
  // deferで必ずFinishを呼ぶ
  defer ctrl.Finish()

  // mock設定。想定される引数と、戻り値がある関数であれば戻り値をReturnメソッドで設定する
  // 設定時に引数へ指定したものと異なる値がテスト時に渡されるとエラーとなり、テストが失敗する
  mock := example.NewMockInterfaceName(ctrl)
  // 2つの引数を受け、2つの戻り値（正常時の値、エラー）を返す関数をmockする場合は以下の様に書く
  mock.EXPECT().mockTargetFunc(paramA, paramB).Return(expectValue, expectError)

  // mock化対象のフィールドを持つ構造体に上記で生成したmockを設定するå
  sample := New(mock)
  sample.exec()
```