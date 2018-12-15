package main

import (
	b64 "encoding/base64"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nlopes/slack"
)

var chanName = []string{"公益"}
var chanCode = []string{"CEVRD231V"}
var slackToken = "xoxp-166092624144-167490026647-504867983490-cce0efd3adb8cc1f520c316fc9eaafb5"

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
	channelEnc := c.Query("name")
	channelbyte, _ := b64.StdEncoding.DecodeString(channelEnc)
	messageEnc := c.Query("message")
	messagebyte, _ := b64.StdEncoding.DecodeString(messageEnc)
	fmt.Println("channle:" + string(channelbyte) + "message:" + string(messagebyte))
	slackSendMessage(identifyChan(string(channelbyte)), string(messagebyte))
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
	channelID, timestamp, err := slackapi.PostMessage(channel.Code, slack.MsgOptionText(msg, false))
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Printf("Message successfully sent to channel %s at %s", channelID, timestamp)
}
