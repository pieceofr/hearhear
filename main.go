package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nlopes/slack"
)

var chanName = []string{"公益"}
var chanCode = []string{"CEVRD231V"}
var slackToken = "xoxp-166092624144-167490026647-504760819587-cc8a485e1f50a2bf078a83fb2cec1dd1"

type slackChan struct {
	Name string
	Code string
}

func main() {
	log.Println("Run HearHear ...")
	runHttpServ("8080")
}

func runHttpServ(port string) {
	route := gin.Default()
	route.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "test ok")
	})

	route.GET("/sendMessage", sendMessage)
	route.Run(":" + port) // listen and serve on 0.0.0.0:8080
}

func test(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "ok",
	})
}

func sendMessage(c *gin.Context) {
	log.Println("start sendMessage")
	channel := c.Param("channel")
	message := c.Param("message")
	slackSendMessage(identifyChan(channel), message)
	c.JSON(200, gin.H{
		"status": "ok",
	})
}

func identifyChan(c string) slackChan {
	socialGoodChan := slackChan{
		Name: chanName[0],
		Code: chanCode[0]}
	return socialGoodChan
}

func slackSendMessage(channel slackChan, msg string) {
	slackapi := slack.New(slackToken)
	log.Println("slack send to " + channel.Name + ":" + msg)
	attachment := slack.Attachment{
		Pretext: "",
		Text:    msg,
		// Uncomment the following part to send a field too
		/*
			Fields: []slack.AttachmentField{
				slack.AttachmentField{
					Title: "a",
					Value: "no",
				},
			},
		*/
	}
	channelID, timestamp, err := slackapi.PostMessage(channel.Code,
		slack.MsgOptionText(msg, false),
		slack.MsgOptionAttachments(attachment))
	//_, _, err := slackapi.PostMessage(channel.Name, msg, slack.PostMessageParameters{})
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	fmt.Printf("Message successfully sent to channel %s at %s", channelID, timestamp)
}
