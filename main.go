package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"github.com/nlopes/slack"
)

var chanName = []string{"公益"}
var chanCode = []string{"CEVRD231V"}
var slackToken = "xoxp-166092624144-167490026647-504798328002-c11e5b02a0790bdbe4a57a5fe610bb4f"

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
	route.Run(":" + port)
	route.POST("/sendMessage", sendMessage)
	route.Run() // listen and serve on 0.0.0.0:8080
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
	_, _, err := slackapi.PostMessage(channel.Name, msg, slack.PostMessageParameters{})
	if err != nil {
		glog.Error(err)
	}

	channelID, timestamp, err := slackapi.PostMessage(channel.Code, msg, slack.PostMessageParameters{})
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Printf("Message successfully sent to channel %s at %s", channelID, timestamp)
}
