package Core

type Message struct {
	Code int
	Msg  string
	Data map[string]interface{}
}

func NewMessage(c int, m string, d map[string]interface{}) Message {

	return Message{
		Code: c,
		Msg:  m,
		Data: d,
	}
}
