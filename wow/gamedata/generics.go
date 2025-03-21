package gamedata

const (
	UP   string = "UP"
	DOWN string = "DOWN"
)

type href struct {
	Href string `json:"href"`
}

type idAndKey struct {
	ID  int  `json:"id"`
	Key href `json:"key"`
}

type idAndType struct {
	ID   int    `json:"id"`
	Type string `json:"type"`
}
