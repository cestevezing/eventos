package response

type GenericResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewGenericResponse(code int, message string, data interface{}) *GenericResponse {
	return &GenericResponse{code, message, data}
}
