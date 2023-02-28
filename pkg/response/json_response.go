package response

type JsonResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewJsonResponse(status int, message string, data interface{}) JsonResponse {
	return JsonResponse{
		Status:  status,
		Message: message,
		Data:    data,
	}
}
