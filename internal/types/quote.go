package types

type Quote struct {
	ID     int64  `json:"id" db:"id"`
	Author string `json:"author" db:"author"`
	Quote  string `json:"quote" db:"quote"`
}
