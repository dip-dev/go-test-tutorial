# Chapter5
### 検証ロジックをassertに書き換えてみよう
chapter1〜chapter4のテストコードの検証ロジックを  
`"github.com/stretchr/testify/assert"`を使って書き換えてください  

assertを利用することによりテストコード内のif文が減り、検証ロジックをシンプルに書くことができます  
また、途中でFailした場合でもテストケースを最後まで実行できるため、どのケースが失敗しているのかを一覧で出すことができます  

※`chapter4/main/sample_test.go`にサンプルがあります  
使い方や詳細は以下を参考に各自で調べてみてください  
https://github.com/stretchr/testify
