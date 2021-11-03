package discord

import (
	"fmt"
	discordgo "github.com/bwmarrin/discordgo"
	"github.com/enescakir/emoji"
	"hash/fnv"
	"log"
	"the-watcher-bot/models"
	"the-watcher-bot/param"
	"time"
)

var (
	DISCORD_TOKEN, CHANNEL_ID, SQUAD_NAME = param.GetDiscordEnvs()
)

var BotId string
var hashInMemory []uint32

func ConnectWithDiscord() *discordgo.Session {
	goBot, err := discordgo.New("Bot " + DISCORD_TOKEN)
	if err != nil {
		fmt.Println(err.Error())
	}
	u, err := goBot.User("@me")
	if err != nil {
		fmt.Println(err.Error())
	}
	BotId = u.ID
	goBot.AddHandler(messageHandler)
	err = goBot.Open()
	if err != nil{
		fmt.Println(err.Error())
	}

	log.Println("Connected with Discord!")

	return goBot

}

func SendAlert(s *discordgo.Session, listOfExceptions []models.Message){
	addOnListOfHashsAndSendToDiscord(s, listOfExceptions)
}


func messageHandler(s *discordgo.Session, m* discordgo.MessageCreate){
	if m.Author.ID == BotId{
		return
	}
	if m.Content == "The Watcher is alive?"{
		_,_ = s.ChannelMessageSend(m.ChannelID, "Yes, I'm alive!")
	}
}

//This function is responsible for checking if the exception is in memory and if it is not added in memory and sends an alert to discord.
func addOnListOfHashsAndSendToDiscord(s *discordgo.Session,message []models.Message){

	listLengh := len(hashInMemory)

	for _, m := range  message {
		hash := generatehash(m.Msg + m.TimeStamp.String())
		_, found := find(hashInMemory,hash)
		if !found {
			s.ChannelMessageSendEmbed(CHANNEL_ID, mountMessage(m))
			hashInMemory = append(hashInMemory, hash)
			time.Sleep(1 * time.Second)
		}
	}

	if listLengh == len(hashInMemory){
		hour,min,sec := time.Now().Clock()
			log.Printf("No new ERROR identified | timeStamp - %v:%v:%v",hour,min,sec)
	}
}


func mountMessage(m models.Message) *discordgo.MessageEmbed{

	emoji := emoji.Robot
	discordMessage := &discordgo.MessageEmbed{
		Title: fmt.Sprintf("ERROR found on application: %v - %v", m.Application, emoji),
		Color:       15158332,
		Description: SQUAD_NAME + ": " + m.Msg,
	}

	return discordMessage
}

/*
func checkBeforeSend(message models.Message){
	newhash := generatehash(message.Msg + message.TimeStamp.String())



}
*/

func find(slice []uint32, val uint32) (int, bool) {
	for i, hash := range slice {
		if hash == val {
			return i, true
		}
	}
	return -1, false
}


func generatehash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}
