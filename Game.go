package main

import (
	"sync"

	"github.com/bwmarrin/discordgo"
)

// Game -
type Game struct {
	Players []string
}

// Games -
var Games = make(map[string]Game)
var mutex = &sync.Mutex{}

// StartGame - function will handle callback for !start, and will add a new gamestate struct to the global map
func StartGame(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "!start" {
		var gameInstance Game
		Games[m.ChannelID] = gameInstance
		s.ChannelMessageSend(m.ChannelID, ":wolf: A new game of Werewolf is starting! For a tutorial, type !help.\r\n\r\n")
	}
}

// JoinGame - callback for when someone says !join
func JoinGame(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "!join" {
		if game, ok := Games[m.ChannelID]; ok {
			mutex.Lock()
			game.Players = append(game.Players, m.Author.Username)
			mutex.Unlock()
			s.ChannelMessageSend(m.ChannelID, m.Author.Username+" has joined the game.")
		} else {
			var gameInstance Game
			mutex.Lock()
			Games[m.ChannelID] = gameInstance
			mutex.Unlock()
			s.ChannelMessageSend(m.ChannelID, ":wolf: A new game of Werewolf is starting! For a tutorial, type !help.\r\n\r\n")
			mutex.Lock()
			game.Players = append(game.Players, m.Author.Username)
			mutex.Unlock()
			s.ChannelMessageSend(m.ChannelID, m.Author.Username+" has joined the game.")
		}
	}
}
