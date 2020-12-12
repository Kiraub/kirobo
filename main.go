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

var subCommands = map[string]uint8{
	"help": help,
	"run":  run,
}

var runFs *flag.FlagSet = flag.NewFlagSet("main", flag.PanicOnError)
var subCommand string
var token string

func init() {
	runFs.StringVar(&token, "t", "", "App token")
}

func main() {
	exeName := os.Args[0]
	if len(os.Args) < 3 {
		logger.Errorf("%v requires a subcommand to execute", exeName)
		return
	}
	subCommand = os.Args[1]
	runFs.Parse(os.Args[2:])
	logger.Debugf("%v, %v", subCommand, token)

	k := kirobo.BuildKirobo()
	err := k.Connect(token)
	if err != nil {
		logger.Errorf("Something unexpected happened: %v", err)
		panic(err)
	}
	k.EnablePingPong(true)
	sc := make(chan os.Signal)
	signal.Notify(sc, os.Interrupt)
	<-sc
	err = k.Disconnect()
	if err != nil {
		logger.Errorf("Something unexpected happened: %v", err)
		panic(err)
	}
}
