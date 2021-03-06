package kirobo

import (
	"kirobo/features"
	"kirobo/logger"
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
)

const (
	// PingPong is a messaging feature, replying to "Ping" with "Pong!" and vice versa
	PingPong features.FeatureKey = iota
)

// Kirobo is a wrapped discord session
type Kirobo struct {
	session  *discordgo.Session
	log      *logger.Logger
	features map[features.FeatureKey]features.Feature
}

// BuildKirobo assembles a new kirobo struct that can than be used to communicate with discord's bot api.
// Built-in features are also registered but remain disabled by default.
func BuildKirobo(logPrefix string) *Kirobo {
	k := new(Kirobo)
	k.session = nil
	k.log = logger.New(os.Stdout, os.Stderr, logPrefix, log.LstdFlags|log.Lshortfile)
	k.features = make(map[features.FeatureKey]features.Feature)
	k.features[PingPong] = features.PingPongFeature(k.log)
	return k
}

// Connect connects to discords bot api using the provided token.
// If a session already exists with the same token, this function does nothing.
// If a session already exists with another token, that session is closed first.
func (r *Kirobo) Connect(token string) error {
	r.log.Enter("Connect")
	defer r.log.Exit("Connect")
	// Close existing session if it has a different token
	if r.session != nil {
		if r.session.Identify.Token == token {
			return nil
		}
		if err := r.session.Close(); err != nil {
			return err
		}
	}
	dgSession, err := discordgo.New("Bot " + token)
	if err != nil {
		return err
	}
	r.session = dgSession
	if err := r.session.Open(); err != nil {
		r.log.Errorf("Error connecting: %v", err)
	}
	r.log.Debugf("Connected %v", r.session.State.SessionID)
	return nil
}

// Disconnect ...
func (r *Kirobo) Disconnect() error {
	r.log.Enter("Disconnect")
	defer r.log.Exit("Disconnect")
	if r.session == nil {
		return nil
	}
	r.log.Debugf("Disconnected")
	return r.session.Close()
}

// ToggleFeature ...
func (r *Kirobo) ToggleFeature(fKey features.FeatureKey, enable bool) (success bool) {
	r.log.Enter("ToggleFeature")
	defer r.log.Exit("ToggleFeature")
	success = true
	r.log.Debugf("Attempting to toggle feature %v", fKey)
	f, ok := r.features[fKey]
	if !ok {
		r.log.Errorf("Cannot toggle unknown feature key %v", fKey)
		success = false
		return
	}
	if enable {
		if f.IsEnabled() {
			r.log.Debugf("Feature %v is already enabled", fKey)
		} else {
			f.Enable(r.session.AddHandler(f.Handler()))
			r.log.Debugf("Successfully enabled feature %v", fKey)
		}
	} else {
		if !f.IsEnabled() {
			r.log.Debugf("Feature %v is already disabled", fKey)
		} else {
			f.Disable()
			r.log.Debugf("Successfully disabled feature %v", fKey)
		}
	}
	return
}
