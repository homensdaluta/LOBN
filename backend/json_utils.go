package main

type Message struct {
	Success any
	Error   Error
}

type Error struct {
	Code    string
	Message string
	Detail  string
}

func SendSuccess(body any) Message {
	data := Message{
		Success: body,
	}
	return data
}

func SendError(code string, message string, detail string, body any) Message {
	data := Message{
		Success: body,
		Error: Error{
			Code:    code,
			Message: message,
			Detail:  detail,
		},
	}
	return data
}
