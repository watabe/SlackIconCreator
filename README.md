SlackIconCreator
====

コマンドラインで簡単にSlack用の文字アイコンを作成するツールです。Golangで作りました。また、ttfを引数に渡すことで好きなフォントを利用してSlack用の文字アイコンを作ることが出来ます。

## Descrpition
このツールでは以下のことが出来ます。
- Slack用の文字アイコンを生成することが出来ます（128*128/透過png）
- 複数行の文字で作成することが出来ます
- 128*128の画像にうまいことサイズを合わせて画像を作成します（中寄せ）

また、今後以下の拡張を想定しています
- 背景色の変更
- 文字色の変更
- 複数行表示の際の `|` をエスケープする

## Usage
以下のようなコマンドを実行することでpngを出力することが出来ます。

```
./SlackIconCreator -ttf TTFフォント -out 出力先ファイル名.png"-mes "あと|よろ"
```

## Args
`-h` を実行すれば引数についてのヘルプが出力されます。 `-mes` に関しては、 `|` で改行を表しているのでご注意ください。

```
-mes string
  	出力したいテキスト。'|'で改行（複数行化/エスケープは未対応）になる
-out string
  	出力先のファイル名
-ttf string
  	利用するTTFファイル
```

## Install
ソースコードをダウンロードして `go build SlackIconCreator.go` すればビルドできます。
必要なライブラリが不足している場合は都度 `go get` してください。

## Licence
[Apache License 2.0](https://github.com/watabe/SlackIconCreator/blob/master/LICENSE)
