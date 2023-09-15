package doc

type Document struct {
	Id     int    `json:"id,omitempty"`
	Type   string `json:"type"`
	Number string `json:"number"`
}

type DocumentList struct {
	Documents []Document `json:"documents"`
}
