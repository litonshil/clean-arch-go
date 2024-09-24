package msgutil

import (
	"fmt"
)

type Data map[string]interface{}

type Msg struct {
	Data Data
}

func NewMessage() Msg {
	return Msg{
		Data: make(Data),
	}
}

func (m Msg) Set(key string, value interface{}) Msg {
	m.Data[key] = value
	return m
}

func (m Msg) Done() Data {
	return m.Data
}

func RequestBodyParseErrorResponseMsg(entity, action string) Data {
	return NewMessage().Set("message", fmt.Sprintf("failed to parse %s %s request body", entity, action)).Done()
}

func RequestQueryParamParseErrorResponseMsg() Data {
	return NewMessage().Set("message", "failed to parse query param").Done()
}

func SomethingWentWrongMsg() Data {
	return NewMessage().Set("message", "something went wrong").Done()
}

func UnprocessableEntityMsg() Data {
	return NewMessage().Set("message", "Unprocessable entity").Done()
}

func InvalidUserRequest() Data {
	return NewMessage().Set("message", "Invalid user request").Done()
}

func NotFoundMsg() Data {
	return NewMessage().Set("message", "not found").Done()
}

func UpdateSuccessMsg(entity string) Data {
	return NewMessage().Set("message", fmt.Sprintf("%v successfully updated", entity)).Done()
}

func ErrorMsg(msg string) Data {
	return NewMessage().Set("message", msg).Done()
}

func ForbiddenResponseMsg() Data {
	return NewMessage().Set("message", "request not granted on this resource").Done()
}

func CustomSuccessMsgWithEntityID(id int, msg Data) interface{} {
	return map[string]interface{}{
		"id":      id,
		"message": msg["message"],
	}
}

func InvalidCredentialsMsg() Data {
	return NewMessage().Set("message", "invalid username or password").Done()
}
