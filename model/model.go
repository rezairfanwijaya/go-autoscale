package model

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Worker struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
