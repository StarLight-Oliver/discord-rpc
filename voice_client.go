package discordrpc

type VoiceSettings struct {
	AutomaticGainControl *bool `json:"automatic_gain_control,omitempty"`
	EchoCancellation     *bool `json:"echo_cancellation,omitempty"`
	NoiseSuppression     *bool `json:"noise_suppression,omitempty"`
	QualityOfService     *bool `json:"quality_of_service,omitempty"`
	SilenceWarning       *bool `json:"silence_warning,omitempty"`
	Deaf                 *bool `json:"deaf,omitempty"`
	Mute                 *bool `json:"mute,omitempty"`
}

func (vs *VoiceSettings) SetAutomaticGainControl(b bool) {
	vs.AutomaticGainControl = &b
}

func (vs *VoiceSettings) SetEchoCancellation(b bool) {
	vs.EchoCancellation = &b
}

func (vs *VoiceSettings) SetNoiseSuppression(b bool) {
	vs.NoiseSuppression = &b
}

func (vs *VoiceSettings) SetQualityOfService(b bool) {
	vs.QualityOfService = &b
}

func (vs *VoiceSettings) SetSilenceWarning(b bool) {
	vs.SilenceWarning = &b
}

func (vs *VoiceSettings) SetDeaf(b bool) {
	vs.Deaf = &b
}

func (vs *VoiceSettings) SetMute(b bool) {
	vs.Mute = &b
}

func (client *Client) SetVoiceSettings(activity *VoiceSettings) error {
	return client.sendCommandWithInterface(SetVoiceSettingsCommand, activity)
}
