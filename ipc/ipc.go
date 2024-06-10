package ipc

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
)

type Socket struct {
	net.Conn
}

func (s *Socket) Send(opcode int, data string) (string, error) {
	buf := new(bytes.Buffer)

	err := binary.Write(buf, binary.LittleEndian, int32(opcode))
	if err != nil {
		return "", err
	}

	err = binary.Write(buf, binary.LittleEndian, int32(len(data)))
	if err != nil {
		return "", err
	}

	buf.Write([]byte(data))

	_, err = s.Write(buf.Bytes())

	if err != nil {
		return "", err
	}

	return s.Read()
}

func (s *Socket) Read() (string, error) {
	buf := make([]byte, 1024)
	dataLen, err := s.Conn.Read(buf)
	if err != nil {
		return "", err
	}

	buffer := new(bytes.Buffer)
	for i := 8; i < dataLen; i++ {
		buffer.WriteByte(buf[i])
	}

	// if dataLength equals 1024, we need to read more
	for dataLen == 1024 {
		dataLen, err = s.Conn.Read(buf)
		if err != nil {
			return "", err
		}

		for i := 0; i < dataLen; i++ {
			buffer.WriteByte(buf[i])
		}
	}

	r := buffer.String()
	if r == "" {
		return "", fmt.Errorf("empty response")
	}

	return r, nil
}
