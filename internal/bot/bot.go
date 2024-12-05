package bot

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"discord-automod/internal/bot/cogs"

	"github.com/bwmarrin/discordgo"
)

var (
	BotToken string = os.Getenv("TOKEN")
)

func Run(sid string, pwd string) {
	discord, err := discordgo.New("Bot " + BotToken)
	if err != nil {
		log.Fatal(err)
		return
	}

	// Add event handler
	discord.AddHandler(cogs.NewMessage)

	// Open session
	discord.Open()
	defer discord.Close()

	// Run until code is terminated
	fmt.Println("Bot running...")
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

}
