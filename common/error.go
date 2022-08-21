package common

type RestErr struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}

func BadRequest(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  400,
		Error:   "Bad request",
	}
}

func Forbidden(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  403,
		Error:   "User with that email already exists",
	}
}

func InternalErr(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  500,
		Error:   "server error",
	}
}

func NotFound(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  404,
		Error:   "User with that id does not exist",
	}
}
