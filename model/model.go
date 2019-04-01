package model

type Url struct {
	ID           string `json:"id,omitempty"`
	To           string `json:"to,omitempty"`
	Description  string `json:"description,omitempty"`
	Creationtime int64  `json:"creationtime,omitempty"`
}
