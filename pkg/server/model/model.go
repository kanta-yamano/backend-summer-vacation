package model

type User struct {
	Name string `json:"name"`
}
type DateTime struct {
	Timestamp string `json:"timestamp"`
	Details   Detail `json:"detail"`
}
type Detail struct {
	Date string `json:"date"`
	Time string `json:"time"`
}
type Zeller struct {
	Year  int `json:"year"`
	Month int `json:"month"`
	Day   int `json:"day"`
}
type Weeks struct {
	Week string `json:"week"`
}
type SignUp struct {
	Id       string `json:id`
	Password string `json:password`
}

type Restoken struct {
	Token string `json:"token"`
}
