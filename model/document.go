package model

type Document struct {
	Id     int
	Type   string
	Number string
}

type DocumentList struct {
	Documents []Document
}
