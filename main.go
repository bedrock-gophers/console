package main

import (
	"github.com/bedrock-gophers/console/console"
	"github.com/df-mc/dragonfly/server"
	"github.com/df-mc/dragonfly/server/player/chat"
	"log"
	"log/slog"
)

func main() {
	// Enable the console.
	console.Enable(slog.Default())
	chat.Global.Subscribe(chat.StdoutSubscriber{})

	conf, err := server.DefaultConfig().Config(slog.Default())
	if err != nil {
		log.Fatalln(err)
	}

	srv := conf.New()
	srv.CloseOnProgramEnd()

	srv.Listen()
	for range srv.Accept() {

	}
}
