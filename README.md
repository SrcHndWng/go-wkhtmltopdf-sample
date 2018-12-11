# go-wkhtmltopdf-sample

## About This

https://github.com/SebastiaanKlippert/go-wkhtmltopdf/blob/master/simplesample_test.go

上記のソースを参考に

* divタグによる罫線の出力
* 日本語を出力
* テンプレートでのHTML定義
* 画像の読み込み

をやってみたもの。

画像についてはsample.tpl内のimgタグに任意のURL等を設定すること。

## タグ

* v1.0.0
    * 固定のHTML形式文字列を出力。
* v2.0.0
    * テンプレートでHTMLを定義
    * 画像を出力
    * 日本語を出力
