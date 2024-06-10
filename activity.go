package discordrpc

import (
	"os"
	"strconv"
	"time"
)

type Activity struct {
	Details    string      `json:"details,omitempty"`
	State      string      `json:"state,omitempty"`
	Assets     *Assets     `json:"assets,omitempty"`
	Timestamps *TimeStamps `json:"timestamps,omitempty"`
	Party      *Party      `json:"party,omitempty"`
	Secrets    *Secrets    `json:"secrets,omitempty"`

	Instance bool `json:"instance,omitempty"`
}

type Assets struct {
	LargeImage string `json:"large_image,omitempty"`
	LargeText  string `json:"large_text,omitempty"`
	SmallImage string `json:"small_image,omitempty"`
	SmallText  string `json:"small_text,omitempty"`
}

type TimeStamps struct {
	Start *Epoch `json:"start,omitempty"`
	End   *Epoch `json:"end,omitempty"`
}

type Party struct {
	ID   string `json:"id,omitempty"`
	Size []int  `json:"size,omitempty"`
}

type Secrets struct {
	Join     string `json:"join,omitempty"`
	Spectate string `json:"spectate,omitempty"`
	Match    string `json:"match,omitempty"`
}

type Epoch struct{ time.Time }

func (t Epoch) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(t.Unix(), 10)), nil
}

func (client *Client) SetActivity(activity *Activity) error {

	payload := make(map[string]interface{})

	payload["pid"] = os.Getpid()
	payload["activity"] = activity

	_, err := client.sendCommand(SetActivityCommand, &payload)

	return err
}
