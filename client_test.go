package discordrpc

import (
	"fmt"
	"time"
)

func Example() {
	client, err := NewClient("client id", "")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer client.Socket.Close()

	fmt.Println("Connected to Discord RPC")

	err = client.SetActivity(&Activity{
		Details: "Details",
		State:   "State",
		Assets: &Assets{
			LargeImage: "keyart_hero",
			SmallImage: "keyart_hero",
		},
		Timestamps: &TimeStamps{
			Start: &Epoch{Time: time.Now()},
		},
	})

	if err != nil {
		fmt.Println(err)
		return
	}
}
