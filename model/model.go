package model

type Url struct {
	ID           string `json:"id,omitempty"`
	To           string `json:"to,omitempty"`
	Description  string `json:"description,omitempty"`
	// Creator      string `json:"creator,omitempty"`
	// Creationtime string `json:"creationtime,omitempty"`
}

/*type User struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}*/
