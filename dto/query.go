package dto

type QueryBody struct {
	Queries []string `json:"queries"`
}

type QueryResponse struct {
	Data  [][]map[string]interface{} `json:"data"`
	Error *string                    `json:"error"`
}
