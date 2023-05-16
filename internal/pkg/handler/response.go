package handler

type ResponseBody struct {
	Status   string      `json:"status"`
	Message  string      `json:"message"`
	Data     interface{} `json:"data,omitempty"`
	MetaData interface{} `json:"metadata,omitempty"`
}
