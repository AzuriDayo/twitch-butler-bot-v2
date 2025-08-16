package main

import (
	"log"
	"os"

	"github.com/slazurin/twitch-butler-bot-v2/pkg/twitchbutlerbot"
)

func main() {
	instance, err := twitchbutlerbot.NewInstance(twitchbutlerbot.InstanceProps{
		BroadcasterAuth: &twitchbutlerbot.InstancePropsTwitchAuth{
			Token:        os.Getenv("TWITCH_BROADCASTER_TOKEN"),
			RefreshToken: os.Getenv("TWITCH_BROADCASTER_REFRESH_TOKEN"),
			ChannelName:  os.Getenv("TWITCH_BROADCASTER_CHANNEL_NAME"),
			ClientID:     os.Getenv("TWITCH_BROADCASTER_CLIENT_ID"),
			ClientSecret: os.Getenv("TWITCH_BROADCASTER_CLIENT_SECRET"),
		},
		BotAuth: twitchbutlerbot.InstancePropsTwitchAuth{
			Token:        os.Getenv("TWITCH_BOT_TOKEN"),
			RefreshToken: os.Getenv("TWITCH_BOT_REFRESH_TOKEN"),
			ChannelName:  os.Getenv("TWITCH_BOT_CHANNEL_NAME"),
			ClientID:     os.Getenv("TWITCH_BOT_CLIENT_ID"),
			ClientSecret: os.Getenv("TWITCH_BOT_CLIENT_SECRET"),
		},
		AutoReconnect: true,
	})
	if err != nil {
		log.Fatalln(err)
	}

	// Group these in a helper function later on
	// Build constructs for main modules, such as attendance, custom commands, spotify song requests
	instance.AddChannelRewardRedemptionHandler()
	// instance.AddChannelRewardRedemptionHandler()
	// instance.AddChannelRewardRedemptionHandler()
	// instance.AddChannelRewardRedemptionHandler()

	instance.AddCommandHandler()
	// instance.AddCommandHandler()
	// instance.AddCommandHandler()
	// instance.AddCommandHandler()

	log.Fatalln(instance.Run())
}
