package model

type Trigger struct {
	ID     int    `json:"id"`
	User   string `json:"user"`
	Device string `json:"device"`
}
