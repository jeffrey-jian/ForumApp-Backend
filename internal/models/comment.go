package models

type Comment struct {
	ID           int    `json:"id"`
	Post_ID      int    `json:"post_id"`
	Author_ID    int    `json:"author_id"`
	Date_created string `json:"date_created"`
	Comment_text string `json:"comment_text"`
}
