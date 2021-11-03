package main

import (
	"context"
	"log"
	"the-watcher-bot/discord"
	"the-watcher-bot/es"
	"the-watcher-bot/param"
	"time"
)


func main(){
	log.Println("I'm the watchman and I'm alive")

	collectInteval := param.GetGeneralEnvs()
	interval, _ := time.ParseDuration(collectInteval)


	ctx := context.Background()

	client, _ := es.Connection()
	discordSession := discord.ConnectWithDiscord()

	ticker := time.NewTicker(interval)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <- ticker.C:
				listOfExceptions := es.GetElasticInfo(client,ctx)
				discord.SendAlert(discordSession, listOfExceptions)
			case <- quit:
				ticker.Stop()
				return
			}
		}

	}()


	<-make(chan  struct{})

}

