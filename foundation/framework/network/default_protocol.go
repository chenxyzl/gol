package network

import (
	"encoding/binary"
	"io"
)

type DefaultPacket struct {
	buff []byte
}

func (p *DefaultPacket) Serialize() []byte {
	return p.buff
}

func (p *DefaultPacket) GetBody() []byte {
	return p.buff[4:]
}

func NewDefaultPacket(buff []byte) *DefaultPacket {
	p := &DefaultPacket{}

	p.buff = make([]byte, 4+len(buff))
	binary.BigEndian.PutUint32(p.buff[0:4], uint32(len(buff)))
	copy(p.buff[4:], buff)

	return p
}

type DefaultProtocol struct {
}

func (p *DefaultProtocol) ReadPacket(r io.Reader) (Packet, error) {
	var (
		lengthBytes = make([]byte, 4)
		length      uint32
	)

	// read length
	if _, err := io.ReadFull(r, lengthBytes); err != nil {
		return nil, err
	}
	length = binary.BigEndian.Uint32(lengthBytes)
	//if length > 1024 {
	//	return nil, errors.New("the size of packet is larger than the limit")
	//}

	buff := make([]byte, length)

	// read body ( buff = lengthBytes + body )
	if _, err := io.ReadFull(r, buff); err != nil {
		return nil, err
	}

	return NewDefaultPacket(buff), nil
}
