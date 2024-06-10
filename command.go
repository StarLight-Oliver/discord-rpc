package discordrpc

const (
	// DispatchCommand event dispatch
	DispatchCommand command = "DISPATCH"

	// AuthorizeCommand used to authorize a new client with your app
	AuthorizeCommand command = "AUTHORIZE"

	// AuthenticateCommand used to authenticate an existing client with your app
	AuthenticateCommand command = "AUTHENTICATE"

	// GetGuildCommand used to retrieve guild information from the client
	GetGuildCommand command = "GET_GUILD"

	// GetGuildsCommand used to retrieve a list of guilds from the client
	GetGuildsCommand command = "GET_GUILDS"

	// GetChannelCommand used to retrieve channel information from the client
	GetChannelCommand command = "GET_CHANNEL"

	// GetChannelsCommand used to retrieve a list of channels for a guild from the client
	GetChannelsCommand command = "GET_CHANNELS"

	// SubscribeCommand used to subscribe to an RPC event
	SubscribeCommand command = "SUBSCRIBE"

	// UnSubscribeCommand used to unsubscribe from an RPC event
	UnSubscribeCommand command = "UNSUBSCRIBE"

	// SetUserVoiceSettingsCommand used to change voice settings of users in voice channels
	SetUserVoiceSettingsCommand command = "SET_USER_VOICE_SETTINGS"

	// SelectVoiceChannelCommand used to join or leave a voice channel, group dm, or dm
	SelectVoiceChannelCommand command = "SELECT_VOICE_CHANNEL"

	// GetSelectedVoiceChannelCommand used to get the current voice channel the client is in
	GetSelectedVoiceChannelCommand command = "GET_SELECTED_VOICE_CHANNEL"

	// SelectTextChannelCommand used to join or leave a text channel, group dm, or dm
	SelectTextChannelCommand command = "SELECT_TEXT_CHANNEL"

	// GetVoiceSettingsCommand used to retrieve the client's voice settings
	GetVoiceSettingsCommand command = "GET_VOICE_SETTINGS"

	// SetVoiceSettingsCommand used to set the client's voice settings
	SetVoiceSettingsCommand command = "SET_VOICE_SETTINGS"

	// CaptureShortcutCommand used to capture a keyboard shortcut entered by the user
	CaptureShortcutCommand command = "CAPTURE_SHORTCUT"

	// SetCertifiedDevicesCommand used to send info about certified hardware devices
	SetCertifiedDevicesCommand command = "SET_CERTIFIED_DEVICES"

	// SetActivityCommand used to update a user's Rich Presence
	SetActivityCommand command = "SET_ACTIVITY"

	// SendActivityJoinInviteCommand used to consent to a Rich Presence Ask to Join request
	SendActivityJoinInviteCommand command = "SEND_ACTIVITY_JOIN_INVITE"

	// CloseActivityRequestCommand used to reject a Rich Presence Ask to Join request
	CloseActivityRequestCommand command = "CLOSE_ACTIVITY_REQUEST"
)
