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
	PhotoId      uint64
	UserId       uint64
	Text_Comment string
}

type PhotoList struct {
	PList []Photo
}
