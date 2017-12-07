package main

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/bwmarrin/discordgo"
)

type Game struct {
	PlayerMap map[string]*PlayerAtt
}

type PlayerAtt struct {
	Role int
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
			Games[m.ChannelID] = &Game{make(map[string]*PlayerAtt)}
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
		if Games[m.ChannelID] == nil {
			Games[m.ChannelID] = &Game{make(map[string]*PlayerAtt)}
			s.ChannelMessageSend(m.ChannelID, ":wolf: A new game of Werewolf is starting! For a tutorial, type !help.\r\n\r\n")
			Games[m.ChannelID].PlayerMap[m.Author.Username] = &PlayerAtt{Role: 0}
			s.ChannelMessageSend(m.ChannelID, m.Author.Username+" has joined the game!")
		} else {
			Games[m.ChannelID].PlayerMap[m.Author.Username] = &PlayerAtt{Role: 0}
			s.ChannelMessageSend(m.ChannelID, m.Author.Username+" has joined the game!")
		}
	}
}

func LeaveGame(s *discordgo.Session, m *discordgo.MessageCreate) {
	mutex.Lock()
	defer mutex.Unlock()

	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "!leave" {
		if game, ok := Games[m.ChannelID]; ok {
			fmt.Println(game)
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
		if Games[m.ChannelID] != nil {
			for k, v := range Games[m.ChannelID].PlayerMap {
				fmt.Println("Player: " + k + " Role: " + strconv.Itoa(v.Role))
				s.ChannelMessageSend(m.ChannelID, "Player: "+k+" Role: "+strconv.Itoa(v.Role))
			}

		}
	}
}
