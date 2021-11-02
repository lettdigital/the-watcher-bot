package param

import (
	"log"
	"os"
)

func GetElasticEnvs() (string, string, string, string,string,string) {

	host := getOrPanic("ELASTIC_HOST")
	index := getOrPanic("INDEX")
	applications := getOrPanic("APPLICATIONS")
	targets := getOrPanic("TARGETS")
	level := getOrPanic("LEVEL")
	queryRange := getOrPanic("QUERY_RANGE")

	return host, index, applications, targets,level,queryRange
}


func GetDiscordEnvs() (string, string, string) {

	discordToken := getOrPanic("DISCORD_TOKEN")
	channelId := getOrPanic("CHANNEL_ID")
	squadName := getOrPanic("SQUAD_NAME")


	return discordToken, channelId, squadName
}

func GetGeneralEnvs() string{
	 return getOrPanic("COLLECT_INTERVAL")
}


// getOrPanic gets the specified environment variable or logs an error and exit
func getOrPanic(env string) string {
	value := os.Getenv(env)

	if value == "" {
		log.Panicf("Error loading envioment:: %v", value)
	}
	return value
}

// sltop string literal to pointer
func sltop(lit string) *string {
	return &lit
}
