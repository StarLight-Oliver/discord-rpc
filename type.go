package discordrpc

type handshake struct {
	Version  string `json:"v"`
	ClientID string `json:"client_id"`
}

type ResponseData struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
