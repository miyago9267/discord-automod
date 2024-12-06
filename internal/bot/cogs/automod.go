package cogs

import (
	"discord-automod/internal/models"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

var muteDuration = 10 * 60 // 10 minutes
var filter *models.WordFilter

func InitFilter(filePath string) error {
	var err error
	filter, err = models.NewWordFilter(filePath)
	if err != nil {
		return err
	}
	log.Println("Filter initialized successfully.")
	return nil
}

func NewMessage(discord *discordgo.Session, message *discordgo.MessageCreate) {
	userID := message.Author.ID

	if userID == discord.State.User.ID {
		return
	}

	content := strings.TrimSpace(message.Content)

	// Check if the message is a command
	if strings.HasPrefix(content, "!addban ") {
		handleAddBanCommand(discord, message, content)
		return
	}

	// Check if the message is a command
	if strings.HasPrefix(content, "!delban ") {
		handleDelBanCommand(discord, message, content)
		return
	}

	// Iterate all Banned Word in filter
	words := strings.Fields(strings.ToLower(message.Content))
	for _, word := range words {
		if filter.IsBanned(word) {
			err := muteUser(discord, message.GuildID, userID)
			if err != nil {
				log.Printf("Failed to mute user %s: %v", userID, err)
			} else {
				log.Printf("Muted user %s for %v", userID, muteDuration)
				discord.ChannelMessageSend(message.ChannelID, fmt.Sprintf("用戶 <@%s> 已被禁言 %v，因為多次違反伺服器守則", userID, muteDuration))
			}
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

func handleAddBanCommand(session *discordgo.Session, message *discordgo.MessageCreate, content string) {
	parts := strings.SplitN(content, " ", 2)
	if len(parts) != 2 {
		session.ChannelMessageSend(message.ChannelID, "用法錯誤，請使用 `!addban <word>`")
		return
	}

	word := strings.TrimSpace(parts[1])
	if word == "" {
		session.ChannelMessageSend(message.ChannelID, "請提供有效的禁用字詞。")
		return
	}

	err := filter.AddBannedWord(word)
	if err != nil {
		session.ChannelMessageSend(message.ChannelID, "無法新增禁用字詞："+err.Error())
		return
	}

	session.ChannelMessageSend(message.ChannelID, "成功新增禁用字詞："+word)
}

func handleDelBanCommand(session *discordgo.Session, message *discordgo.MessageCreate, content string) {
	parts := strings.SplitN(content, " ", 2)
	if len(parts) != 2 {
		session.ChannelMessageSend(message.ChannelID, "用法錯誤，請使用 `!delban <word>`")
		return
	}
	word := strings.TrimSpace(parts[1])
	if word == "" {
		session.ChannelMessageSend(message.ChannelID, "請提供有效的禁用字詞。")
		return
	}

	err := filter.DeleteBannedWord(word)
	if err != nil {
		session.ChannelMessageSend(message.ChannelID, "無法刪除禁用字詞："+err.Error())
		return
	}

	session.ChannelMessageSend(message.ChannelID, "成功刪除禁用字詞："+word)
}
