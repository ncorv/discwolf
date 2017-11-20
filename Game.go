package main

import (
	"strings"
	"sync"

	"github.com/bwmarrin/discordgo"
)

type Game struct {
	Players []string
}

var Games = make(map[string]*Game)
var mutex = &sync.Mutex{}

func StartGame(s *discordgo.Session, m *discordgo.MessageCreate) {
	mutex.Lock()
	defer mutex.Unlock()

	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "!start" {
		Games[m.ChannelID] = &Game{}
		s.ChannelMessageSend(m.ChannelID, ":wolf: A new game of Werewolf is starting! For a tutorial, type !help.\r\n\r\n")
	}
}

func JoinGame(s *discordgo.Session, m *discordgo.MessageCreate) {
	mutex.Lock()
	defer mutex.Unlock()

	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "!join" {
		if game, ok := Games[m.ChannelID]; ok {
			game.Players = append(game.Players, m.Author.Username)
			s.ChannelMessageSend(m.ChannelID, m.Author.Username+" has joined the game!")

		} else {
			Games[m.ChannelID] = &Game{}
			s.ChannelMessageSend(m.ChannelID, ":wolf: A new game of Werewolf is starting! For a tutorial, type !help.\r\n\r\n")
			game.Players = append(game.Players, m.Author.Username)
			s.ChannelMessageSend(m.ChannelID, m.Author.Username+" has joined the game!")
		}
	}
}

func PrintGame(s *discordgo.Session, m *discordgo.MessageCreate) {
	mutex.Lock()
	defer mutex.Unlock()

	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "!PrintGame" {
		if game, ok := Games[m.ChannelID]; ok {
			s.ChannelMessageSend(m.ChannelID, "Players: "+strings.Join(game.Players, ", "))
		}
	}
}
