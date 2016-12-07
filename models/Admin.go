package models

func GetAdmin() *Admin {
	var admin = new(Admin)
	return admin
}

type Admin struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
