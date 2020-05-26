package models

type Input struct {
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Data        map[string]string `json:"data"`
}
