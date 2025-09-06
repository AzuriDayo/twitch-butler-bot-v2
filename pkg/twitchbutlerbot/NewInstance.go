package twitchbutlerbot

import (
	"errors"
	"log"

	"github.com/nicklaw5/helix/v2"
)

func buildHelixClient(auth InstancePropsTwitchAuth) (*helix.Client, error) {
	client, err := helix.NewClient(&helix.Options{
		ClientID:        auth.ClientID,
		ClientSecret:    auth.ClientSecret,
		UserAccessToken: auth.Token,
		RefreshToken:    auth.RefreshToken,
	})
	return client, err
}

func NewInstance(props InstanceProps) (*instance, error) {
	i := &instance{}
	// validate props
	hasBroadcasterAuth := props.BroadcasterAuth != nil

	// configure event sub connections
	if hasBroadcasterAuth {
		broadcasterHelixClient, err := buildHelixClient(*props.BroadcasterAuth)
		if err != nil {
			return nil, errors.New("NewInstance: Failed to create broadcasterHelixClient " + err.Error())
		}
		i.broadcasterHelixClient = broadcasterHelixClient

		broadcasterSelf, err := i.broadcasterHelixClient.GetUsers(&helix.UsersParams{})
		if len(broadcasterSelf.Data.Users) < 1 || err != nil {
			return nil, errors.New("NewInstance: Failed to query broadcasterSelf " + err.Error())
		}
		log.Println("Preflight broadcaster", broadcasterSelf.Data.Users[0].Login, "success")

		i.broadcasterChannelName = broadcasterSelf.Data.Users[0].Login

	}
	botHelixClient, err := buildHelixClient(props.BotAuth)
	if err != nil {
		return nil, errors.New("NewInstance: Failed to create botHelixClient " + err.Error())
	}
	i.botHelixClient = botHelixClient

	botSelf, err := i.botHelixClient.GetUsers(&helix.UsersParams{})
	if len(botSelf.Data.Users) < 1 || err != nil {
		return nil, errors.New("NewInstance: Failed to query botSelf " + err.Error())
	}
	log.Println("Preflight bot", botSelf.Data.Users[0].Login, "success")

	i.botChannelName = botSelf.Data.Users[0].Login

	return i, nil
}
