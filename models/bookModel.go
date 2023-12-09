package models

import "time"

type Book struct {
	Name string `form:"name"`
	Author string `form:"author_name"`
	PublishedDate string `form:"publish_date"`
	Image string `form:"image_url"`
	Location string `form:"location"` 
}

type Person struct {
	Name     string `form:"name"`
	Address  string `form:"address"`
	Birthday time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
}

type PersonBindURI struct {
	ID string `uri:"id" binding:"required,uuid"`
	Name string `uri:"name" binding:"required"`
}

