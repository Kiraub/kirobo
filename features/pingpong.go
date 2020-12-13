package features

import (
	"kirobo/logger"
	"regexp"
	"strings"

	"github.com/bwmarrin/discordgo"
)

// PingPongFeature ...
func PingPongFeature(l *logger.Logger) Feature {
	pingPong := func(session *discordgo.Session, message *discordgo.MessageCreate) {
		l.Debugf("Handling MessageCreate event: %v", message.Message.Content)
		if message.Author.ID == session.State.User.ID {
			l.Debugf("Same user as bot, ignoring")
			return
		}
		var messageContent = message.Message.Content
		if ping, err := regexp.MatchString("^ping$", strings.ToLower(messageContent)); err != nil {
			if ping {
				session.ChannelMessageSend(message.ChannelID, "Pong!")
			} else if pong, err := regexp.MatchString("^pong$", strings.ToLower(messageContent)); err != nil {
				if pong {
					session.ChannelMessageSend(message.ChannelID, "Ping!")
				}
			} else {
				l.Errorf(err.Error())
			}
		} else {
			l.Errorf(err.Error())
		}
	}
	ppf := CreateFeature(pingPong)
	return ppf
}
