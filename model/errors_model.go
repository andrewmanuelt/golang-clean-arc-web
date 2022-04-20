package model

type Errors struct {
	Message string `json:"message"`
	Errors  interface{}
}
