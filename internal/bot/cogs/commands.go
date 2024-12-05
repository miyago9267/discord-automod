package cogs

import "github.com/bwmarrin/discordgo"

func NewMessage(discord *discordgo.Session, message *discordgo.MessageCreate) {
	// Ignore bot messaage
	if message.Author.ID == discord.State.User.ID {
		return
	}
}
