package response

type Failed struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func FailedWithData(data interface{}, message string) Failed {
	return Failed{Status: "failed", Data: data, Message: message}
}

func FailedMessage(message string) Failed {
	return Failed{Status: "failed", Message: message}
}
