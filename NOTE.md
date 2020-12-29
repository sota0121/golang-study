# Goプログラミングノート

## ランタイム環境構築方法

- 公式サイトからバイナリ（dmg）を取得してインストール
- アンインストールする方法は[こちら](https://golang.org/doc/manage-install#uninstalling)
  - ただし、1.13以前の場合は、`/usr/local/bin/go`にインストールされるので、注意

## 開発環境構築

- ビジネス用途でがっつりやるなら有料だけど[Goland](https://www.jetbrains.com/ja-jp/go/)が一番良さそう
- ただし、ここでは無料で使えるVSCodeを前提とする

---

- Go用エクステンション `Go from Go team at Google` をインストールすればOK
- その他の支援系パッケージはVSCodeに任せた

## 実行方法

- local golang runtime
  - build
    - `go build xxx.go` then xxx made as single bin exec file
    - `go build .`
    - `go build PKG_NAME`
  - build and run
    - `go run xxx.go`
    - `go run .`
    - `go run PKG_NAME`
  - run only
    - `./BIN_NAME`
- Playground at local
  - use [gp](https://github.com/tenntenn/goplayground) cmd
- Playground at browser
  - [Golang PlayGround](https://play.golang.org/)

## コーディングノート

- 自前パッケージを作るときは…
  - `main.go`のディレクトリに、パッケージ名のディレクトリ(pkgxとする)を作り、その中にパッケージ用のコード（ファイル名はなんでもいい）を保存し、各ファイルのパッケージ名指定部分（`package pkgx`）をパッケージ名にする
  - 自作関数を他のファイルから呼び出したいときは、「エクスポート用の関数名」にしないといけない。そのルールは、「頭文字が大文字であること」

## 課題ノート

### Gopher道場1

- Goはどのような特徴の言語か？
  - Googleが開発したOSS言語
  - 静的型付け言語、シンプルな構文、言語レベルのlint機能
  - 並行処理が書きやすい、シングルバイナリ、クロスコンパイルが容易、標準ライブラリが豊富
- どのようなソフトウェアを開発するのに向いているか考えてみましょう。
  - 
- Goはどのような背景で開発されたか？
  - 開発速度を上げるため、マルチコア時代のシステムプログラミング言語として、LLライクな言語、静的解析しやすい言語
- Goを開発に用いることで、自身の開発環境のどういう問題が解決しそうかか考えてみましょう。

## Gopher道場4：パッケージ

- TRY: [パッケージを分ける](https://docs.google.com/presentation/d/1jrXP4ggIWYSbC9WQhBgccm51jwS3gdizMYAc0OLL6s0/edit#slide=id.g8ba2c4ad64_0_937)
  - 
