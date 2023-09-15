package client

type Pong struct {
	Pong string `json:"pong"`
}

type Error struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

type Success struct {
	Ok bool `json:"ok"`
}
