package models

type Like struct {
	ID          int `json:"id"`
	Post_ID     int `json:"post_id"`
	Liked_by_ID int `json:"liked_by_id"`
}
