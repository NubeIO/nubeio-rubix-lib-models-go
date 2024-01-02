package dto

type Token struct {
	Name    string `json:"name" binding:"required"`
	Token   string `json:"token" binding:"required"`
	Blocked *bool  `json:"blocked"`
}

type TokenCreate struct {
	Name    string `json:"name" binding:"required"`
	Blocked *bool  `json:"blocked" binding:"required"`
}

type TokenBlock struct {
	Blocked *bool `json:"blocked" binding:"required"`
}
