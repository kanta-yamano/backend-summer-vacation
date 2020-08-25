package model

type User struct {
	Name string `json:"name"`
}
type Detail struct {
	datetime
}
type datetime struct {
	Date string `json:"date"`
	Time string `json:"time"`
}
