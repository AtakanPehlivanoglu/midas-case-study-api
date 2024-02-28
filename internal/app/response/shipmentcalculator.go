package response

type ShipmentCalculatorResponse struct {
	HTTPStatusCode int `json:"-"` // http response status code

	Message string `json:"message"` // user-level status message
}

func NewShipmentCalculatorResponse(statusCode int, message string) *ShipmentCalculatorResponse {
	return &ShipmentCalculatorResponse{
		HTTPStatusCode: statusCode,
		Message:        message,
	}
}
