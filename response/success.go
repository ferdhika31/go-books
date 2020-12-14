package response

type Success struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func OKWithData(data interface{}, message string) Success {
	return Success{Status: "ok", Data: data, Message: message}
}

func OKMessage(message string) Success {
	return Success{Status: "ok", Message: message}
}
