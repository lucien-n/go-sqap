package models

type TodoRequest struct {
	Content    string `json:"content"`
	IsComplete bool   `json:"is_complete"`
}

type Todo struct {
	UID        string `json:"uid"`
	Content    string `json:"content"`
	IsComplete bool   `json:"is_complete"`
}
