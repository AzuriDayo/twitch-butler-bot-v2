package twitchbutlerbot

type InstancePropsTwitchAuth struct {
	Token        string
	RefreshToken string
	ChannelName  string
	ClientID     string
	ClientSecret string
}

type InstanceProps struct {
	BotAuth         InstancePropsTwitchAuth
	BroadcasterAuth *InstancePropsTwitchAuth
	AutoReconnect   bool
}
