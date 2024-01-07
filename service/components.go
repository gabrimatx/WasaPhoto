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
	Id            uint64 `json:"id"`
	PhotoId       uint64 `json:"photoId"`
	PublisherId   uint64 `json:"publisherId"`
	PublisherName string `json:"publisherName"`
	CommentText   string `json:"text"`
}

type CommentList struct {
	CList []CommentListElement
}

type Response struct {
	PhotoList
	UserName      string `json:"userName"`
	FollowCount   int    `json:"followCount"`
	FollowedCount int    `json:"followedCount"`
	IsFollowed    bool   `json:"isFollowed"`
	IsBanned      bool   `json:"isBanned"`
}

type PhotoListElement struct {
	Id          uint64 `json:"id"`
	ReleaseDate string `json:"date"`
	Caption     string `json:"caption"`
	PublisherId uint64 `json:"userId"`
	Likes       int    `json:"likecount"`
}

type PhotoStreamListElement struct {
	Id            uint64 `json:"id"`
	ReleaseDate   string `json:"date"`
	Caption       string `json:"caption"`
	PublisherName string `json:"publisherName"`
	Likes         int    `json:"likecount"`
}

type PhotoList struct {
	PList []PhotoListElement
}

type PhotoStreamList struct {
	PList []PhotoStreamListElement
}
