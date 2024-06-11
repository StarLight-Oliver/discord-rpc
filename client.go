package discordrpc

import (
	"encoding/json"
	"fmt"

	"github.com/StarLight-Oliver/discord-rpc/ipc"
	"github.com/google/uuid"
)

type Client struct {
	ClientID     string
	ClientSecret string
	Socket       *ipc.Socket
}

type command string

func NewClient(clientID string, clientSecret string) (*Client, error) {

	if clientID == "" {
		return nil, fmt.Errorf("clientID is required")
	}

	socket, err := ipc.NewSocket()
	if err != nil {
		return nil, err
	}

	handShake := &handshake{
		Version:  "1",
		ClientID: clientID,
	}

	packetData, err := json.Marshal(handShake)
	if err != nil {
		return nil, err
	}

	client := &Client{clientID, clientSecret, socket}

	response, err := client.Socket.Send(0, string(packetData))

	if err != nil {
		return nil, err
	}

	var responseData ResponseData
	err = json.Unmarshal([]byte(response), &responseData)
	if err != nil {
		return nil, err
	}

	if responseData.Code > 1000 {
		return nil, fmt.Errorf("handshake error: %s", responseData.Message)
	}

	return client, nil
}

type payload struct {
	Cmd   command                 `json:"cmd"`
	Args  *map[string]interface{} `json:"args"`
	Event string                  `json:"evt,omitempty"`
	Data  *ResponseData           `json:"data,omitempty"`
	Nonce uuid.UUID               `json:"nonce"`
}

type payloadInterface struct {
	Cmd   command       `json:"cmd"`
	Args  interface{}   `json:"args"`
	Event string        `json:"evt,omitempty"`
	Data  *ResponseData `json:"data,omitempty"`
	Nonce uuid.UUID     `json:"nonce"`
}

func (client *Client) sendCommand(command command, data *map[string]interface{}) (string, error) {

	nonce := uuid.New()

	payloadRaw := &payload{
		Cmd:   command,
		Nonce: nonce,
		Args:  data,
	}

	packetData, err := json.Marshal(payloadRaw)
	if err != nil {
		return "", err
	}

	response, err := client.Socket.Send(1, string(packetData))
	if err != nil {
		return "", err
	}

	var responseData payload
	err = json.Unmarshal([]byte(response), &responseData)
	if err != nil {

		if err.Error() == "json: cannot unmarshal string into Go struct field ResponseData.data.code of type int" {

			if command == AuthorizeCommand {
				return response, nil
			}
		}

		return "", err
	}

	if responseData.Data.Code > 1000 {
		return "", fmt.Errorf("command error: %s", responseData.Data.Message)
	}

	if responseData.Nonce != nonce {
		return "", fmt.Errorf("invalid nonce")
	}

	return response, nil
}

func (client *Client) sendCommandWithInterface(command command, data interface{}) error {

	nonce := uuid.New()

	payloadRaw := &payloadInterface{
		Cmd:   command,
		Nonce: nonce,
		Args:  data,
	}

	packetData, err := json.Marshal(payloadRaw)
	if err != nil {
		return err
	}

	response, err := client.Socket.Send(1, string(packetData))
	if err != nil {
		return err
	}

	var responseData payload
	err = json.Unmarshal([]byte(response), &responseData)
	if err != nil {
		return err
	}

	if responseData.Data.Code > 1000 {
		return fmt.Errorf("command error: %s", responseData.Data.Message)
	}

	if responseData.Nonce != nonce {
		return fmt.Errorf("invalid nonce")
	}

	return nil
}
