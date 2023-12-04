package api

type Error struct {
	code    string
	message string
}

type User struct {
	Id   int
	Name string
}

type Photo struct {
	Id          int
	File        string
	ReleaseDate string
	Caption     string
	PublisherId int
	Likes       int
}

type Comment struct {
	Id           int
	PhotoId      int
	UserId       int
	Text_Comment string
}

type PhotoList struct {
	PhotoList []Photo
}
