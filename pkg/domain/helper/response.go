package helper

// Response : custom json response
type Response struct{
	Code int `json:"code"`
	Message string `json:"message"`
	Data map[string]interface{} `json:"data,omitempty"`
}
