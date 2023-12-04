package api

type Error struct {
	code    string
	message string
}

type User struct {
	Id   int
	name string
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
	photoId      int
	userId       int
	text_comment string
}

type PhotoList struct {
	PhotoList []Photo
}
