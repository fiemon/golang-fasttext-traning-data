package main

import (
	"github.com/bluele/mecab-golang"
	"flag"
	"os"
	"log"
	"encoding/csv"
	"io"
	"strings"
	"regexp"
)

var i  = flag.String("i", " string", "help message for \"i\" option")
var o  = flag.String("o", " string", "help message for \"o\" option")

func createTestData(m *mecab.MeCab) {
	tg, err := m.NewTagger()
	if err != nil {
		panic(err)
	}
	defer tg.Destroy()

	inf, err := os.Open(*i)
	if err != nil {
		log.Fatal("failed:input file open")
	}
	defer inf.Close()
	reader := csv.NewReader(inf) //utf8
	//reader.Comma = ','
	reader.LazyQuotes = true
	reader.FieldsPerRecord = -1

	ouf, err := os.Create(*o)
	if err != nil {
		log.Fatal("failed:output file open")
	}
	defer ouf.Close()

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		lt, err := m.NewLattice(record[0])
		if err != nil {
			panic(err)
		}
		r := regexp.MustCompile(`</doc>|<doc id=.*>|\n`)
		if r.MatchString(record[0]) {
			continue
		}
			//</doc>
		//<doc id="10" url="https://ja.wikipedia.org/wiki?curid=10" title="言語">

		defer lt.Destroy()

		s := ""
		node := tg.ParseToNode(lt)
		for {
			features := strings.Split(node.Feature(), ",")
			//if features[0] == "名詞" || features[0] == "動詞" || features[0] == "形容詞" || features[0] == "副詞" {
			if features[0] == "名詞" {
				str := node.Surface()
				if str != "" {
					s = s + " " + str
				}
			}
			if node.Next() != nil {
				break
			}
		}
		//log.Println(s)
		//os.Exit(100)
		ouf.Write(([]byte)(s))
	}
}

func main() {
	flag.Parse()
	m, err := mecab.New("-d /usr/local/lib/mecab/dic/mecab-ipadic-neologd")
	if err != nil {
		panic(err)
	}
	defer m.Destroy()
	createTestData(m)
}