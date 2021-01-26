package procotol

type Message struct {
	Content []byte
}

func (m *Message) Len() int {
	return len(m.Content)
}

func (m *Message) Empty() bool {
	return m.Len() == 0
}