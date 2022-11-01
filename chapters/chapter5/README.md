# Chapter5
### 検証ロジックをassertに書き換えてみよう
chapter1〜chapter4のテストコードの検証ロジックを  
`"github.com/stretchr/testify/assert"`を使って書き換えてください。  

assertを利用することによりテストコード内のif文が減り、検証ロジックをシンプルに書くことができます。  
また、`assert.Equals`等はテストが途中でFailしても後続の処理が続くため  
どこのテストが失敗したかを一覧で出すことができます。

※`chapter4/main/sample_test.go`にサンプルがあります。 
使い方や詳細は各自で調べてみてください。
https://github.com/stretchr/testify
