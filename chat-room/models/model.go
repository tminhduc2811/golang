package models

type User struct {
	Username string 	`json:"username"`
	Password string		`json:"password"`
	HashPassword []byte	`json:"hash_password,omitempty"`
}

type JwtToken struct {
	Token string `json:"token"`
}

type Message struct {
	Username string	`json:"username"`
	Mess     string	`json:"mess"`
}

type Messages [] Message
