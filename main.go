package main

import (
	"flag"
	"kirobo/kirobo"
	"kirobo/logger"
	"os"
	"os/signal"
)

const (
	help uint8 = iota
	run
)

const logPrefix = "MAIN#"

var log *logger.Logger
var runFs *flag.FlagSet

var subCommands = map[string]uint8{
	"help": help,
	"run":  run,
}

var subCommand string
var token string

func init() {
	runFs = flag.NewFlagSet("run", flag.PanicOnError)
	runFs.StringVar(&token, "t", "", "App token")

	log = logger.CreateLogger()
	log.InfoFormat = logPrefix + logger.InfoFormat
	log.DebugFormat = logPrefix + logger.DebugFormat
	log.ErrorFormat = logPrefix + logger.ErrorFormat
}

func main() {
	exeName := os.Args[0]
	if len(os.Args) < 3 {
		log.Errorf("%v requires a subcommand to execute", exeName)
		return
	}
	subCommand = os.Args[1]
	runFs.Parse(os.Args[2:])
	log.Debugf("%v, %v", subCommand, token)

	k := kirobo.BuildKirobo("KIROBO")
	err := k.Connect(token)
	if err != nil {
		log.Errorf("Something unexpected happened: %v", err)
		panic(err)
	}
	k.ToggleFeature(kirobo.PingPong, true)
	sc := make(chan os.Signal)
	signal.Notify(sc, os.Interrupt)
	<-sc
	log.Infof("Interrupt received")
	err = k.Disconnect()
	if err != nil {
		log.Errorf("Something unexpected happened: %v", err)
		panic(err)
	}
}
