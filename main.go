package main

import (
	"bytes"
	"flag"
	"fmt"
	"holo_check_CLI/model"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/api/googleapi/transport"
	"google.golang.org/api/youtube/v3"
)

var CHANNEL_ID_LIST = map[string]string{
	"ときのそら":      "UCp6993wxpyDPHUpavwDFqgg",
	"AZKi":       "UC0TXe_LYZ4scaW2XMyi5_kw",
	"ロボ子さん":      "UCDqI2jOz0weumE8s7paEk6g",
	"さくらみこ":      "UC-hM6YJuNYVAmUWxeIr9FeA",
	"白上フブキ":      "UCdn5BQ06XqgXoAxIhbqw5Rg",
	"夏色まつり":      "UCQ0UDLQCjY0rmuxCDE38FGg",
	"夜空メル":       "UCD8HOxPs4Xvsm8H0ZxXGiBw",
	"赤井はあと":      "UC1CfXB_kRs3C-zaeTG3oGyg",
	"アキ・ローゼンタール": "UCFTLzh12_nrtzqBPsTCqenA",
	"湊あくあ":       "UC1opHUrw8rvnsadT-iGp7Cg",
	"癒月ちょこ":      "UC1suqwovbL1kzsoaZgFZLKg",
	"百鬼あやめ":      "UC7fk0CB07ly8oSl0aqKkqFg",
	"紫咲シオン":      "UCXTpFs_3PqI41qX2d9tL2Rw",
	"大空スバル":      "UCvzGlP9oQwU--Y0r9id_jnA",
	"大神ミオ":       "UCp-5t9SrOQwXMU7iIjQfARg",
	"戌神ころね":      "UChAnqc_AY5_I3Px5dig3X1Q",
	"猫又おかゆ":      "UCvaTdHTWBGv3MKj3KVqJVCw",
	"不知火フレア":     "UCvInZx9h3jC2JzsIzoOebWg",
	"白銀ノエル":      "UCdyqAaZDKHXg4Ahi7VENThQ",
	"宝鐘マリン":      "UCCzUftO8KOVkV4wQG1vkUvg",
	"兎田ぺこら":      "UC1DCedRgGHBdm81E1llLhOQ",
	"潤羽るしあ":      "UCl_gCybOJRIgOXw6Qb4qJzQ",
	"星街すいせい":     "UC5CwaMl1eIgY8h02uZw7u8A",
	"天音かなた":      "UCZlDXzGoo7d44bwdNObFacg",
	"桐生ココ":       "UCS9uQI-jC3DE0L4IpXyvr6w",
	"角巻わため":      "UCqm3BQLlJfvkTsX_hvm0UmA",
	"常闇トワ":       "UC1uv2Oq6kNxgATlCiez59hw",
	"姫森ルーナ":      "UCa9Y57gfeY0Zro_noHRVrnw",
}

func main() {
	// command line arguments
	var (
		target_name = flag.String("name", "", "Name of target channel")
	)
	flag.Parse()

	//.env ファイルからAPI_KEYを取得
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatalf("failed loading .env: %v", err)
		return
	}

	API_KEY := os.Getenv("API_KEY")

	client := &http.Client{
		Transport: &transport.APIKey{Key: API_KEY},
	}

	service, err := youtube.New(client)
	if err != nil {
		log.Fatalf("Error creating new youtube client: %v", err)
	}

	part := []string{"snippet"}
	call := service.Search.List(part).ChannelId(*target_name).MaxResults(3).Order("date")
	resp, err := call.Do()
	if err != nil {
		log.Fatalf("failed query: %v", err)
	}
	if resp != nil {
		for _, item := range resp.Items {
			fmt.Printf("%s: %v\n", *target_name, item.Snippet.LiveBroadcastContent)
			fmt.Printf("\t%v\n", item.Snippet.Title)
		}
	}
}

func PrintSearchResults(searchResults []*model.SearchResult) string {
	var resultTable bytes.Buffer
	resultTable.WriteString("--------------------------------------\n")
	resultTable.WriteString(" | owner name |   title    |  status  |\n")
	resultTable.WriteString(" |------------+------------+----------|\n")
	resultTable.WriteString(" |  湊あくあ  | テスト配信 | upcoming |\n")
	resultTable.WriteString(" |------------+------------+----------|")
	return resultTable.String()
}
