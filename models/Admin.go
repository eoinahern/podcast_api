package models

func GetAdmin() *Admin {
	var admin = new(Admin)
	return admin
}

type Admin struct {
	Name         string `json:"name"`
	Email        string `json:"email"  orm:"pk; unique"`
	Password     string `json:"password"`
	Isloggedin   bool   `json:"isloggedin"`
	IsRegistered bool   `json:"isregistered"`
	AuthKey      string `json:"authkey"`
}

type AdminResp struct {
	Status string `json:"status"`
	Admin  Admin  `json:"admin"`
}

func (a *Admin) TableName() string {
	return "admin"
}
