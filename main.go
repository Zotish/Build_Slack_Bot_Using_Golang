package main

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

func main() {
	os.Setenv("Slack_Bot_Token", "xoxb-5762639277398-5771744463732-MgZf85qnXKTrZuJIUV937h7f")
	os.Setenv("CHANNEL_ID", "C05N6MZJZBR")
	api := slack.New(os.Getenv("Slack_Bot_Token"))
	chennelArr := []string{os.Getenv("CHANNEL_ID")}
	filterArr := []string{"cpp.pdf"}
	for i := 0; i < len(filterArr); i++ {
		params := slack.FileUploadParameters{
			Channels: chennelArr,
			File:     filterArr[i],
		}
		file, err := api.UploadFile(params)
		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}
		fmt.Printf("Name : %s ,URL : %s \n", file.Name, file.URL)
	}
}
