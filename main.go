package main

import (
	"fmt"
	"log"
	"net/http"

	"google.golang.org/api/googleapi/transport"
	"google.golang.org/api/youtube/v3"
)

const API_KEY = "AIzaSyDk4VHVENR1IedMhACYJZjJRyf_KMRj11g"
const COCO_CHANNEL_ID = "UCS9uQI-jC3DE0L4IpXyvr6w"
const OKAYU_CHANNEL_ID = "UCvaTdHTWBGv3MKj3KVqJVCw"
const UNKO_CHANNEL_ID = "UCx1nAvtVDIsaGmCMSe8ofsQ"

func main() {
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
