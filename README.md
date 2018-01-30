# 環境構築
<pre>
# MeCab Install
$ brew search mecab
$ brew install mecab mecab-ipadic
$ export CGO_LDFLAGS="`mecab-config --libs`"
$ export CGO_CFLAGS="-I`mecab-config --inc-dir`"
$ brew install mecab mecab-ipadic git curl xz
$ git clone --depth 1 https://github.com/neologd/mecab-ipadic-neologd.git
$ cd mecab-ipadic-neologd
$ ./bin/install-mecab-ipadic-neologd -n -a
# test
$ mecab -d /usr/local/lib/mecab/dic/mecab-ipadic-neologd
# golang package install
$ go get github.com/bluele/mecab-golang
</pre>

# 実行コマンド
<pre>
go run mecab.go -i "input file path" -o "output filepath"
go run testdata/mecabtestdata.go -i "input file path" -o "output filepath"
</pre>