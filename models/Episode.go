package models

import "time"

func GetEpisode() *Episode {
	var eps = new(Episode)
	return eps
}

type Episode struct {
	Id       string    `json:"id" orm:"pk,auto"`
	Email    string    `json:"email"`
	AudioUrl string    `json:"audiourl"`
	Created  time.Time `json:"created" orm:"auto_now_add;type(date)"`
}

func (a *Episode) TableName() string {
	return "episode"
}
