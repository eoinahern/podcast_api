package models

func GetEpisode() *Episode {
	var eps = new(Episode)
	return eps
}

type Episode struct {
	Id       string `json:"id"`
	AudioUrl string `json:"audiourl"`
	Created  string `json:"created"`
}
