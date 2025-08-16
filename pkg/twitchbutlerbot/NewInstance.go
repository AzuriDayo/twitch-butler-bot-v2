package twitchbutlerbot

import (
	"errors"

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
		i.broadcasterChannelName = props.BroadcasterAuth.ChannelName

	}
	botHelixClient, err := buildHelixClient(props.BotAuth)
	if err != nil {
		return nil, errors.New("NewInstance: Failed to create botHelixClient " + err.Error())
	}
	i.botHelixClient = botHelixClient
	i.botChannelName = props.BotAuth.ChannelName

	// configure internal modules if any (ie: custom commands, and reward redemption handlers)

	return i, nil
}
