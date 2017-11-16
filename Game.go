package main

import "github.com/bwmarrin/discordgo"

// StartGame - function will handle callback for !start, and will add a new gamestate struct to the global slice
func StartGame(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "!start" { // and players have all identified and readied{
		s.ChannelMessageSend(m.ChannelID, ":wolf: A new game of Werewolf is starting! For a tutorial, type !help.\r\n\r\n")
	}
}

// JoinGame - callback for when someone says !join
func JoinGame(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "!join" {
		// first add person to list

		// then say they joined the game
		s.ChannelMessageSend(m.ChannelID, m.Author.Username+" has joined the game.")
	}
}

// function to print player list to channel before start
