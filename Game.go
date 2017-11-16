package main

import "github.com/bwmarrin/discordgo"

// Game -
type Game struct {
	Players []string
}

// Games -
var Games map[string]Game = make(map[string]Game)

// StartGame - function will handle callback for !start, and will add a new gamestate struct to the global map
func StartGame(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "!start" { // and players have all identified and readied{
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
		//do we currently have an entry with the same ChannelID
		if Games[m.ChannelID] == nil {
			var gameInstance Game
			Games[m.ChannelID] = gameInstance
			s.ChannelMessageSend(m.ChannelID, ":wolf: A new game of Werewolf is starting! For a tutorial, type !help.\r\n\r\n")
		} else {
			Games[m.ChannelID].Players = append(Games[m.ChannelID].Players, m.Author.Username)
			s.ChannelMessageSend(m.ChannelID, m.Author.Username+" has joined the game.")
		}
	}
}
