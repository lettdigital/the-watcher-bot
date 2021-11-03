	package es

	import (
		"context"
		"encoding/json"
		elastic "github.com/olivere/elastic/v7"
		"log"
		"strings"
		"the-watcher-bot/models"
		"the-watcher-bot/param"
	)

	var (
		HOST, INDEX, APPLICATIONS, TARGETS,LEVEL,QUERYRANGE  = param.GetElasticEnvs()
	)

	func Connection() (*elastic.Client, error) {
		client, err := elastic.NewClient(elastic.SetURL(HOST),
			elastic.SetSniff(false),
			elastic.SetHealthcheck(false))

		if err != nil {
			log.Panic("Error trying to connect with elastic\n", err)
		}

		log.Println("Connected with the elastic!")

		return client, err
	}


	func GetElasticInfo(client *elastic.Client, ctx context.Context)[]models.Message {
		var listOfException []models.Message

		log.Println("Collecting logs")

		applications := APPLICATIONS
		targets := TARGETS

		for _, application := range strings.Split(applications, ",") {

			for _,  target := range strings.Split(targets, ",") {

				applicationMatch := elastic.NewMatchPhraseQuery("application", application)
				targetMatch := elastic.NewMatchPhraseQuery("msg", target)
				level := elastic.NewMatchPhraseQuery("level", LEVEL)


				timeStamp := elastic.NewRangeQuery("@timestamp")
				timeStamp.Gte(QUERYRANGE)
				timeStamp.Lte("now")
				timeStamp.TimeZone("UTC")


				query := elastic.NewBoolQuery().Filter(applicationMatch,targetMatch,level).Filter(timeStamp)

				searchResult, err := client.Search().
					Index(INDEX).
					Query(query).
					From(0).
					Size(10).
					Do(ctx)
				if err != nil {
					log.Printf("[ProductsES][GetPIds]Error=", err)
				}

			for _, hit := range searchResult.Hits.Hits {
				var infoLogs models.InfoLogs
				json.Unmarshal(hit.Source, &infoLogs)
				msg := infoLogs.Msg
				timeStamp := infoLogs.Timestamp
				applicationName:= infoLogs.Application

				 m := models.Message{Msg: msg,TimeStamp: timeStamp,Application: applicationName}
					listOfException = append(listOfException, m)
				}
			}
		}

		return listOfException
	}


