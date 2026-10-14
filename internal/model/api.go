package model

type ListResponse struct {
	Result interface{} `json:"result"`
	Count  int64       `json:"count"`
}
