package components

type Error struct {
	code    string
	message string
}

type User struct {
	Name string
}

type Photo struct {
	Caption     string
	PublisherId int
}

type Comment struct {
	PhotoId      int
	UserId       int
	Text_Comment string
}

type PhotoList struct {
	PList []Photo
}
