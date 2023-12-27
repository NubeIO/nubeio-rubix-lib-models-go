package dto

type BulkErrorResponse struct {
	UUID  *string `json:"uuid,omitempty"`
	Name  *string `json:"name,omitempty"`
	Error *string `json:"error,omitempty"`
}

type BulkResponse struct {
	Data   interface{}          `json:"data"`
	Errors []*BulkErrorResponse `json:"errors"`
}
