package procotol

import (
	"bufio"
	"encoding/binary"
	"io"
)

const HeaderLen = 4

func Decode(rr *bufio.Reader) (*Message, error) {
	headBuff := make([]byte, HeaderLen)
	_, err := io.ReadFull(rr, headBuff)
	if err != nil {
		return nil, err
	}
	bodyLen := binary.BigEndian.Uint32(headBuff)
	bodyBuff := make([]byte, bodyLen)
	_, err = io.ReadFull(rr, bodyBuff)
	if err != nil {
		return nil, err
	}
	return &Message{Content: bodyBuff}, nil
}

func Encode(msg *Message) []byte {
	bodyLen := msg.Len()
	data := make([]byte, HeaderLen+bodyLen)
	binary.BigEndian.PutUint32(data[:4], uint32(bodyLen))
	copy(data[4:], msg.Content)
	return data
}
