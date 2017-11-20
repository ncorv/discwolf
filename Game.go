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
		if Games[m.ChannelID] != nil {
			s.ChannelMessageSend(m.ChannelID, "A game has already been started in this channel...")
		} else {
			Games[m.ChannelID] = &Game{}
			s.ChannelMessageSend(m.ChannelID, ":wolf: A new game of Werewolf is starting! For a tutorial, type !help.\r\n\r\n")
		}
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
			if !sliceContainsString(game.Players, m.Author.Username) {
				game.Players = append(game.Players, m.Author.Username)
				s.ChannelMessageSend(m.ChannelID, m.Author.Username+" has joined the game!")
			} else {
				s.ChannelMessageSend(m.ChannelID, m.Author.Username+", you have already joined the game...")
			}

		} else {
			Games[m.ChannelID] = &Game{}
			s.ChannelMessageSend(m.ChannelID, ":wolf: A new game of Werewolf is starting! For a tutorial, type !help.\r\n\r\n")
			Games[m.ChannelID].Players = append(Games[m.ChannelID].Players, m.Author.Username)
			s.ChannelMessageSend(m.ChannelID, m.Author.Username+" has joined the game!")
		}
	}
}

func sliceContainsString(stringSlice []string, searchString string) bool {
	for _, value := range stringSlice {
		if value == searchString {
			return true
		}
	}
	return false
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
