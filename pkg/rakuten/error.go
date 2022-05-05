package rakuten

type ErrorResponse struct {
	Content     string `json:"error"`
	Description string `json:"error_description"`
}
