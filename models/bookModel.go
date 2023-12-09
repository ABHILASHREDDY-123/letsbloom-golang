package models


type Book struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Author string `json:"author_name"`
	PublishedDate string `json:"publish_date"`
	Image string `json:"image_url"`
	Location string `json:"location"` 
}

