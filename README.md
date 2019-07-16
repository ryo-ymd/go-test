### Go/Gin/gormのテスト
もし手元で動かしてみたいのであれば以下の環境を整える
- goenv(homebrew使用推奨)
- go 1.12.6

このレポジトリをGOPATH配下にクローンする
- GoはGOPATH以下にソース、依存先、成果物を置くらしい
- goenvを使用すると、 `~/go` がGOPATHとなるっぽい
- `GOPATH/src` 以下にソースコードは置く
- パッケージ順になるようにする。またインポートの関係上 `github.com` のパスを使いたい(相対パスを用いたインポートはしない)
- つまり `GOPATH/src/github.com/ryo-ymd/go-test` にクローンしてもらえば動く
- この辺は `go get` の仕組みを少し調べるとなるほどとなる
