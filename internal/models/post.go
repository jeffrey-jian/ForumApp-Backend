package models

type Post struct {
	ID           int    `json:"id"`
	Author_ID    int    `json:"author_id"`
	Date_created string `json:"date_created"`
	Category     string `json:"category"`
	Title        string `json:"title"`
	Post_text    string `json:"post_text"`
}
