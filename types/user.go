package types

type User struct {
	Name         string `json:"name"`
	ID           string `json:"id"`
	Email        string `json:"email"`
	PasswordHash string `json:"_"`
}

type CreateUser struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
