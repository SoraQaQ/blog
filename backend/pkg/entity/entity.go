package entity

type Article struct {
	Id         int64
	Title      string
	Summary    string
	ContentUrl string
	Status     int64
	ViewCount  int64
	Tags       string
	ImageUrl   []string
}

type User struct {
	Id       uint64
	Username string
	Nickname string
	Password string
	Email    string
}
