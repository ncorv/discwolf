package main

import (
	"discwolf/command"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
)

var (
	// Token - discord bot token
	Token string
)

func init() {
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}

func main() {
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("Could not create DiscordGo session because: ,", err)
		return
	}

	dg.AddHandler(messageCreate)
	err = dg.Open()
	if err != nil {
		fmt.Println("Could not create connection because: ", err)
		return
	}

	currentTime := time.Now().Local()
	fmt.Println(currentTime, ": Starting discwolf awooooo.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	dg.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "!help" {
		command.HelpCommand(s, m)
	}
	if m.Content == "!join" {
		JoinGame(s, m)
	}
	if m.Content == "!start" {
		StartGame(s, m)
	}
	if m.Content == "!PrintGame" {
		PrintGame(s, m)
	}
	if m.Content == "!leave" {
	}
	if m.Content == "!end" {
	}
	if m.Content == "!alive" {
	}
	if m.Content == "!status" {
	}
}
