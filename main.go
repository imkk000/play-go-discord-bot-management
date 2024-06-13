package main

import (
	"os"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.TimeOnly})

	client, err := discordgo.New("Bot " + token)
	logErr(err, "new session")
	defer logErr(client.Close(), "close session")

	channels, err := client.GuildChannels(guildID)
	logErr(err, "get channels")
	for _, channel := range channels {
		channel, err := client.Channel(channel.ID)
		logErr(err, "get channel")
		log.Debug().
			Str("ID", channel.ID).
			Str("name", channel.Name).
			Msg("get channel")
	}
}

func logErr(err error, msg string) {
	if err != nil {
		log.Fatal().Err(err).Msg(msg)
	}
}

const (
	token   = ""
	guildID = ""
)
