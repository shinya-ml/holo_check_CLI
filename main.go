package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/api/googleapi/transport"
	"google.golang.org/api/youtube/v3"
)

const COCO_CHANNEL_ID = "UCS9uQI-jC3DE0L4IpXyvr6w"
const OKAYU_CHANNEL_ID = "UCvaTdHTWBGv3MKj3KVqJVCw"
const UNKO_CHANNEL_ID = "UCx1nAvtVDIsaGmCMSe8ofsQ"

func main() {
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
	// query := "channelId=UCx1nAvtVDIsaGmCMSe8ofsQ"
	// query := "The Chemical Brothers - Go"
	call := service.Search.List(part).ChannelId(UNKO_CHANNEL_ID).MaxResults(5).Order("date")

	resp, _ := call.Do()
	for _, item := range resp.Items {
		fmt.Println(item.Snippet.LiveBroadcastContent)
	}

}
