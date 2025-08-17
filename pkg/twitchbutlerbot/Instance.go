package twitchbutlerbot

import (
	"github.com/nicklaw5/helix/v2"
)

type instance struct {
	botHelixClient         *helix.Client
	broadcasterHelixClient *helix.Client
	broadcasterChannelName string
	botChannelName         string
	// eventSubChan      *chan (struct{})
}

func (i *instance) AddCommandHandler() {
	// props: command name, handlerFunc

}

func (i *instance) AddChannelRewardRedemptionHandler() {
	// props: reward redepmtion name, handlerFunc

}

func (i *instance) Run() error {

	return nil
}
