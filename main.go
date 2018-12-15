package main

import (
	b64 "encoding/base64"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nlopes/slack"
	"github.com/spf13/viper"
)

var chanName = []string{"公益"}
var chanCode = []string{"CEVRD231V"}
var userName = []string{"piceofr", "Croc"}
var userCode = []string{"U4XEE0SK1", "UDQ27LXB5"}
var slackToken = "xoxp-166092624144-167490026647-504153669040-cbc5d57bc9c48303fdcbcdd7218d12bd"

type slackChan struct {
	Name string
	Code string
}

func main() {
	log.Println("Run HearHear ...")
	slackToken = InitConfig()
	runHttpServ("8080")
}

func runHttpServ(port string) {
	route := gin.Default()
	route.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "test ok")
	})

	route.GET("/sendMessage", sendMessage)
	route.GET("/sendUser", sendUser)
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
	slackSendMessage(string(messagebyte), "channel", identifyChan(string(channelbyte)))
	c.JSON(200, gin.H{
		"status": "ok",
	})
}

func sendUser(c *gin.Context) {
	userEnc := c.Query("user")
	userbyte, _ := b64.StdEncoding.DecodeString(userEnc)
	messageEnc := c.Query("message")
	messagebyte, _ := b64.StdEncoding.DecodeString(messageEnc)
	slackSendMessage(string(messagebyte), "user", identifyUser(string(userbyte)))
}

func identifyChan(c string) slackChan {
	socialGoodChan := slackChan{
		Name: chanName[0],
		Code: chanCode[0]}
	return socialGoodChan
}

func identifyUser(c string) string {
	return userCode[1]
}

func slackSendMessage(msg string, slacktype string, addition interface{}) {
	slackapi := slack.New(slackToken)

	switch slacktype {
	case "channel":
		channel := addition.(slackChan)
		channelID, timestamp, err := slackapi.PostMessage(channel.Code, slack.MsgOptionText(msg, false))
		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}
		log.Println("slack send to " + channel.Name + ":" + msg)
		fmt.Printf("Message successfully sent to channel %s at %s", channelID, timestamp)
	case "user":
		userID := addition.(string)
		fmt.Println("user id:" + userID)
		_, _, channelID, err := slackapi.OpenIMChannel(userID)

		if err != nil {
			fmt.Printf("%s\n", err)
		}
		slackapi.PostMessage(channelID, slack.MsgOptionText(msg, false))
	}

}
func InitConfig() string {
	viper.SetConfigName("conf")
	viper.AddConfigPath("./")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil { // Handle errors reading the config file
		fmt.Printf("config err:%v \n", err)
	}
	token := viper.GetString("tk")
	return token
}
