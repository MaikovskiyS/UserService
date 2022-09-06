package models

type User struct {
	Id      uint   `json:"id"`
	Name    string `json:"name"`
	Age     uint   `json:"age"`
	Friends []User `json:"friends"`
}

type Friendship struct {
	UserId1 uint `json:"id1"`
	UserId2 uint `json:"id2"`
}
