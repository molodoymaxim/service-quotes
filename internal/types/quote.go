package types

type Quote struct {
	ID     int64  `json:"id" db:"id"`
	Author string `json:"author" db:"author"`
	Text   string `json:"text" db:"quote"`
}
