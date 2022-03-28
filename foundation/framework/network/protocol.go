package network

import (
	"io"
)

type Packet interface {
	Serialize() []byte
}

type Protocol interface {
	ReadPacket(conn io.Reader) (Packet, error)
}
