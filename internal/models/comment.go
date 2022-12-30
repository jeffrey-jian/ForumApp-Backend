package models

type Comment struct {
	ID           int    `json:"id"`
	Author_ID    int    `json:"author_id"`
	Date_created string `json:"date_created"`
	Comment_text string `json:"comment_text"`
	Post_ID      int    `json:"post_id"`
}
