package entities

type HttpErrs struct {
	ErrCode int    `json:"err-code"`
	ErrMsg  string `json:"error-msg"`
}

type HttpResponse struct {
	Data       interface{} `json:"data"`
	Message    string      `json:"message"`
	StatusCode int         `json:"statusCode"`
}
