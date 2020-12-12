package kirobo

import (
	"kirobo/logger"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var discord, err = discordgo.New("Bot " + "authentication token")

// Kirobo is a wrapped discord session
type Kirobo struct {
	dgSession *discordgo.Session
}

// BuildKirobo assembles a new kirobo struct that can than be used to communicate with discord's bot api
func BuildKirobo() *Kirobo {
	k := new(Kirobo)
	k.dgSession = nil
	return k
}

// Connect connects to discords bot api using the provided token.
// If a session already exists with the same token, this function does nothing.
// If a session already exists with another token, that session is closed first.
func (r *Kirobo) Connect(token string) error {
	// Close existing session if it has a different token
	if r.dgSession != nil {
		if r.dgSession.Identify.Token == token {
			return nil
		}
		if err := r.dgSession.Close(); err != nil {
			return err
		}
	}
	dgSession, err := discordgo.New("Bot " + token)
	if err != nil {
		return err
	}
	r.dgSession = dgSession
	if err := r.dgSession.Open(); err != nil {
		logger.Errorf("Error connecting: %v", err)
	}
	logger.Debugf("Connected %v", r.dgSession.State.SessionID)
	return nil
}

// Disconnect ...
func (r *Kirobo) Disconnect() error {
	if r.dgSession == nil {
		return nil
	}
	logger.Debugf("Disconnected")
	return r.dgSession.Close()
}

// EnablePingPong ...
func (r *Kirobo) EnablePingPong(enable bool) error {
	r.dgSession.AddHandler(r.pingPong)
	return nil
}

func (r *Kirobo) pingPong(dg *discordgo.Session, message *discordgo.MessageCreate) {
	logger.Debugf("Handling MessageCreate event: %v", message.Message.Content)
	if message.Author.ID == r.dgSession.State.User.ID {
		logger.Debugf("Same user as bot, ignoring")
		return
	}
	var messageContent = message.Message.Content
	if strings.Contains(strings.ToLower(messageContent), "ping") {
		r.dgSession.ChannelMessageSend(message.ChannelID, "Pong!")
	}
}
