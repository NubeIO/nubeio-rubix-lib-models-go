package dto

type PaginationResponse struct {
	Total  int64       `json:"total"`
	Offset *int        `json:"offset"`
	Limit  *int        `json:"limit"`
	Data   interface{} `json:"data"`
}
