package cogs

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

var muteDuration = 10 * 60 // 10 minutes

func NewMessage(discord *discordgo.Session, message *discordgo.MessageCreate) {
	userID := message.Author.ID

	if userID == discord.State.User.ID {
		return
	}

	if strings.Contains(message.Content, "@everyone") {
		err := muteUser(discord, message.GuildID, userID)
		if err != nil {
			log.Printf("Failed to mute user %s: %v", userID, err)
		} else {
			log.Printf("Muted user %s for %v", userID, muteDuration)
			discord.ChannelMessageSend(message.ChannelID, fmt.Sprintf("用戶 <@%s> 已被禁言 %v，因為多次使用 @everyone！", userID, muteDuration))
		}
	}
}

func muteUser(discord *discordgo.Session, guildID, userID string) error {
	muteEnd := time.Now().Add(time.Duration(muteDuration) * time.Second)

	_, err := discord.GuildMemberEdit(guildID, userID, &discordgo.GuildMemberParams{
		CommunicationDisabledUntil: &muteEnd,
	})
	if err != nil {
		return fmt.Errorf("failed to mute user: %v", err)
	}
	return nil
}
