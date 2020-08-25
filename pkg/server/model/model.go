package model

type User struct {
	Name string `json:"name"`
}
type DateTime struct {
	Timestamp string  `json:"timestamp"`
	Details   *Detail `json:"detail"`
}
type Detail struct {
	Date string `json:"date"`
	Time string `json:"time"`
}
