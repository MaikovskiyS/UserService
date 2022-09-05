package models

type User struct {
	Id      uint     `json:"id"`
	Name    string   `json:"name"`
	Age     uint     `json:"age"`
	Friends []string `json:"friends"`
}

