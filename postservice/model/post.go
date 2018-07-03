package model

type Post struct {
	Id string `json:"id"`
	Title string `json:"title"`
	Content string `json:"content"`
	UserId string `json:"userId"`

	ServedBy string `json:"servedBy""`
}
