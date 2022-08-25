package models

type Message struct {
	Code int
	Data interface{}
	Msg  string
}

func (m *Message) Success(data interface{}, msg string) {
	m.Code = 200
	m.Data = data
	m.Msg = msg
}

func (m *Message) Fail(data interface{}, msg string) {
	m.Code = 500
	m.Data = data
	m.Msg = msg
}
