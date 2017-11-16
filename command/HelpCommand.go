package command

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
)

// HelpCommand - callback for when someone says !help
func HelpCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	var helpMsg string
	helpMsg += "\r\n*How to Play #Werewolf*\r\n------------------------\r\n"
	helpMsg += "Werewolf is a party game of social deduction. Players are private messaged their role when the game begins. \r\n\r\n"
	helpMsg += "_If you are a Villager_, you must find out who the werewolves are based on their voting and your social deduction skills.\r\n"
	helpMsg += "_If you are a Werewolf_, you must pretend you are not a werewolf by lying as best as you can.\r\n\r\n"
	helpMsg += "The game takes place over several Days and Nights. Each Day all players vote on a player to eliminate. The player with the most votes is eliminated. If there is a tie, nobody is lynched. \r\n"
	helpMsg += "_Each night_, the werewolves will be allowed to vote privately on one player to eliminate. The decision must be unanimous. If its not, you'll keep voting until it is. The bot will private message you.\r\n"
	helpMsg += "The villagers win if they eliminate all the werewolves. The werewolves win if they equal or outnumber the remaining players.\r\n\r\n"
	helpMsg += "\r\n"

	var gameCommandsMsg string
	gameCommandsMsg += "*Game Commands*\r\n------------------------\r\n"
	gameCommandsMsg += "`!new` - Create a new lobby for players to !join for the next game\r\n"
	gameCommandsMsg += "`!join` - Join the lobby for the next game\r\n"
	gameCommandsMsg += "`!leave` - Leave the lobby for the next game\r\n"
	gameCommandsMsg += "`!start` - Start the game, when called with no parameters the lobby players are used\r\n"
	gameCommandsMsg += "`!start all` - Starts a new game with everyone in the channel participating\r\n"
	gameCommandsMsg += "`!start @user1 @user2 @user3` - Starts a new game with the specified users participating\r\n"
	gameCommandsMsg += "`!end` - Cause the game to end prematurely\r\n"
	gameCommandsMsg += "`!option` - View or change options.  Use without any parameters for help and current values.\r\n"
	gameCommandsMsg += "`!remindme` - Remind you of your role in the current game\r\n"
	gameCommandsMsg += "`!dead` - Show dead players\r\n"
	gameCommandsMsg += "`!alive` - Show living players\r\n"
	gameCommandsMsg += "`!status` - Show game status\r\n"

	if m.Content == "!help" {
		currentTime := time.Now().Local()
		fmt.Println(currentTime, " : ", m.ChannelID, m.Author.Username, " has asked for help.")
		s.ChannelMessageSend(m.ChannelID, helpMsg)
		s.ChannelMessageSend(m.ChannelID, gameCommandsMsg)
	}
}
