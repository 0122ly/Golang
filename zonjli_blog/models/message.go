package models

type Message struct {
	Code int
	Data interface{}
	Msg  string
}

func (m *Message) Success(data interface{}) {
	m.Code = 200
	m.Data = data
}

func (m *Message) Fail(code int, msg string) {
	m.Code = 500
	m.Msg = msg
}
