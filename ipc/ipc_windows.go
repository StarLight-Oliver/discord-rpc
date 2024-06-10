//go:build windows
// +build windows

package ipc

import (
	"time"

	npipe "gopkg.in/natefinch/npipe.v2"
)

func NewSocket() (*Socket, error) {
	conn, err := npipe.DialTimeout(`\\.\pipe\discord-ipc-0`, time.Second*2)
	if err != nil {
		return nil, err
	}

	return &Socket{conn}, nil
}
