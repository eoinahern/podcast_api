package models

func GetPodcast() *Podcast {
	var podcast = new(Podcast)
	return podcast
}

type Podcast struct {
	Id       int64  `json:"id" orm:"pk, auto, unique"`
	Name     string `json:"name" orm:"pk"`
	Category string `json:"category"`
	WebSite  string `json:"website"`
	Email    string `json:"email" orm:"rel(fk)"`
}

func (a *Podcast) TableName() string {
	return "podcast"
}
