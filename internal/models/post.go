package models

type Post struct {
	ID              int    `json:"id"`
	Author_ID       int    `json:"author_id"`
	Author_Username string `json:"author_username"`
	Category        string `json:"category"`
	Date_created    string `json:"date_created"`
	Title           string `json:"title"`
	Post_text       string `json:"post_text"`
}
