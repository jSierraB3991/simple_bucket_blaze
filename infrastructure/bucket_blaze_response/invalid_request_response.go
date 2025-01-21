package bucketblazeresponse

type InvalidRequestPayResponse struct {
	CodeStatus  uint   `json:"status"`
	CodeError   string `json:"code"`
	Description string `json:"description"`
}
