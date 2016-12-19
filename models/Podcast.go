package models

func GetPodcast() *Podcast {
	var podcast = new(Podcast)
	return podcast
}

type Podcast struct {
	Name     string `json:"name"`
	Category string `json:"category"`
	WebSite  string `json:"website"`
	Email    string `json:"email"`
}

func (a *Podcast) TableName() string {
	return "podcast"
}
