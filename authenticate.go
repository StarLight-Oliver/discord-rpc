package discordrpc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type AuthorizeResponse struct {
	Code string `json:"code"`
}

type AuthorizeResponsePacket struct {
	Data AuthorizeResponse `json:"data"`
}

const DiscordOAuthTokenEndpoint = "https://discord.com/api/oauth2/token"

type OAuthTokenRequest struct {
	GrantType   string `json:"grant_type"`
	Code        string `json:"code"`
	RedirectURI string `json:"redirect_uri"`
	ClientID    string `json:"client_id"`
}

type OAuthTokenResponse struct {
	AccessToken string `json:"access_token"`
}

func OAuth(client *Client, code string) (string, error) {

	payload := OAuthTokenRequest{
		GrantType:   "authorization_code",
		Code:        code,
		RedirectURI: "http://localhost:3344",
		ClientID:    client.ClientID,
	}

	urlV := url.Values{}

	urlV.Set("grant_type", payload.GrantType)
	urlV.Set("code", payload.Code)
	urlV.Set("redirect_uri", payload.RedirectURI)
	urlV.Set("client_id", payload.ClientID)

	req, err := http.NewRequest("POST", DiscordOAuthTokenEndpoint, bytes.NewBufferString(urlV.Encode()))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	req.SetBasicAuth(client.ClientID, client.ClientSecret)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to make request: %w", err)
	}

	defer resp.Body.Close()

	// var data OAuthTokenResponse
	// err = json.NewDecoder(resp.Body).Decode(&data)
	// if err != nil {
	// 	return "", fmt.Errorf("failed to decode response: %w", err)
	// }

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)

	s := buf.String()

	data := OAuthTokenResponse{}
	err = json.Unmarshal([]byte(s), &data)

	if err != nil {
		return "", fmt.Errorf("failed to decode response: %w", err)
	}

	return data.AccessToken, nil
}

func (client *Client) SetCustomAuthorize(authorize func(*Client, string) (string, error)) {
	client.authorize = authorize
}

func (client *Client) Authorize(scopes []string) (string, error) {

	payload := make(map[string]interface{})
	payload["scopes"] = scopes
	payload["client_id"] = client.ClientID

	data, err := client.sendCommand(AuthorizeCommand, &payload)

	if err != nil {
		return "", err
	}

	dataPacket := AuthorizeResponsePacket{}
	err = json.Unmarshal([]byte(data), &dataPacket)

	if err != nil {
		return "", err
	}

	code := dataPacket.Data.Code
	// allow for custom implementations of this in the future
	// we now send it to the authenticate endpoint
	access_token, err := client.authorize(client, code)

	if err != nil {
		return "", err
	}

	return access_token, nil
}

func (client *Client) Authenticate(code string) error {

	payload := make(map[string]interface{})
	payload["access_token"] = code

	_, err := client.sendCommand(AuthenticateCommand, &payload)

	return err
}
