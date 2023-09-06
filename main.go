package main

import (
	"context"
	"flag"
	"log"

	tgClinet "github.com/nutatio/dump-bot/clients/telegram"
	eventconsumer "github.com/nutatio/dump-bot/consumer/eventConsumer"
	"github.com/nutatio/dump-bot/events/telegram"
	"github.com/nutatio/dump-bot/storage/sqlite"
)

const (
	tgBotHost         = "api.telegram.org"
	fileStoragePath   = "files_storage"
	sqliteStoragePath = "data/sqlite/storage.db"
	batchSize         = 100
)

func main() {
	// s := files.New(fileStoragePath)
	s, err := sqlite.New(sqliteStoragePath)
	if err != nil {
		log.Fatalf("can't connect to storage: %w", err)
	}

	s.Init(context.TODO())

	eventsProcessor := telegram.New(
		tgClinet.New(tgBotHost, mustToken()),
		s,
	)

	log.Print("service started")

	consumer := eventconsumer.New(eventsProcessor, eventsProcessor, batchSize)

	if err := consumer.Start(); err != nil {
		log.Fatal("service is stopped", err)
	}
}

func mustToken() string {
	token := flag.String("tg-bot-token", "", "token for access to telegram bot")
	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}
	return *token
}
