package models

type Message struct {
	Code int
	Data interface{}
	Msg  string
}

func (m *Message) Success() {
	m.Code = 200
	m.Msg = "登陆成功"
}

func (m *Message) Fail() {
	m.Code = 500
	m.Msg = "密码错误"
}
