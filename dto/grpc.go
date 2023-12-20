package dto

type GrpcBody struct {
	Api    string `json:"api,omitempty"`
	Method string `json:"method,omitempty"`
	Body   []byte `json:"body,omitempty"`
	Args   string `json:"args,omitempty"`
}

type GrpcResponse struct {
	Body []byte `json:"body,omitempty"`
}
