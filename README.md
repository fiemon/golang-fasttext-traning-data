# 環境構築
<pre>
# MeCab Install
$ brew search mecab
$ brew install mecab mecab-ipadic
$ export CGO_LDFLAGS="`mecab-config --libs`"
$ export CGO_CFLAGS="-I`mecab-config --inc-dir`"
# golang package install
$ go get github.com/bluele/mecab-golang
</pre>

# 環境構築
<pre>
go run mecab.go -i "input file path" -o "output filepath"
</pre>