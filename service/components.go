package components

type User struct {
	Name string
}

type Photo struct {
	Caption string
}

type Comment struct {
	PhotoId      uint64
	UserId       uint64
	Text_Comment string
}

type CommentListElement struct {
	PublisherId uint64
	commentText string
}

type PhotoListElement struct {
	Id          uint64 `json:"id"`
	ReleaseDate string `json:"date"`
	Caption     string `json:"caption"`
	PublisherId uint64 `json:"userId"`
	Likes       int    `json:"likecount"`
}

type PhotoList struct {
	PList []PhotoListElement
}
