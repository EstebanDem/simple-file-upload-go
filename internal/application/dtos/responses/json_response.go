package responses

type JsonResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}
