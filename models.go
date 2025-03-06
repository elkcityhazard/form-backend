package main

type ErrorMsg struct {
	Code    int    `json:"status_code"`
	Message string `json:"error_message"`
	Data    any    `json:"data"`
}

func NewErrorMsg(code int, msg string, data any) ErrorMsg {
	return ErrorMsg{
		Code:    code,
		Message: msg,
		Data:    data,
	}
}

type EmailAndMessage struct {
	Email       string `json:"email"`
	Message     string `json:"message"`
	PhoneNumber string `json:"phone_number"`
}
