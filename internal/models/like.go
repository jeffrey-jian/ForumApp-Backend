package models

type Like struct {
	ID      int `json:"id"`
	User_ID int `json:"user_id"`
	Post_ID int `json:"post_id"`
}
