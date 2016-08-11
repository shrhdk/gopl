# 練習問題10.3

`fetch http://gopl.io/ch1/helloworld?go-get=1`を使って
この本のサンプルコードをホストしているサービスを調べなさい。
(`go get`からのHTTPリクエストは`go-get`パラメータを含んでいるので、
サーバは通常のブラウザのリクエストと区別することができます。)

## 結果

```
$ fetch 'http://gopl.io/ch1/heloworld?go-get=1'
<!DOCTYPE html>
<html>
<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
<meta name="go-import" content="gopl.io git https://github.com/adonovan/gopl.io">
</head>
<body>
</body>
```
