package api

type Error struct {
	code    string
	message string
}

type User struct {
	ID   int
	name string
}

type Photo struct {
	ID          int
	file        string
	releaseDate string
	caption     string
	publisherId int
	likes       int
}

type Comment struct {
	ID           int
	text_comment string
}
