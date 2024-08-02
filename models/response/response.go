package response

type Response struct {
	Result  string      `json:"result`
	Error   bool        `json:"error"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
