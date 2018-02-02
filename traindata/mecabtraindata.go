package traindata

import (
	"github.com/bluele/mecab-golang"
	"flag"
	"os"
	"log"
	"encoding/csv"
	"io"
	"strings"
)

var i  = flag.String("i", " string", "help message for \"i\" option")
var o  = flag.String("o", " string", "help message for \"o\" option")
var l = "__label__"

func createTrainData(m *mecab.MeCab) {
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
	reader.Comma = ','
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
		defer lt.Destroy()

		s := ""
		node := tg.ParseToNode(lt)
		for {
			features := strings.Split(node.Feature(), ",")
			if features[0] == "名詞" || features[0] == "動詞" || features[0] == "形容詞" || features[0] == "副詞" {
				//rep := regexp.MustCompile(`([\s].*|[\t].*|[A-Za-z].*|[0-9].*|[０-９].*|【.+】|\.|!|！|？|~|\?|-|【|】|〟|〝|「|」|=|・|%|\[|\]|&|;|\(|\)|：|◆|#|♂|“)`)
				str := node.Surface()
				//str = rep.ReplaceAllString(str, "")
				if str != "" {
					s = s + " " + str
				}
			}
			if node.Next() != nil {
				break
			}
		}
		c := record[1]
		s = l + c + s + "\n"
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
	createTrainData(m)
}