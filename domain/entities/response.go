package entities

type ResposeMessage struct {
	Message string `json:"message,omitempty"`
	Status  int    `json:"status,omitempty"`
}

type ResposeError struct {
	Message string `json:"message,omitempty"`
	Status  int    `json:"status,omitempty"`
}
