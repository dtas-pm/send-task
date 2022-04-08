package send

import "time"

type Student struct {
	FullName  string   `json:"full_name"`
	Login     string   `json:"login"`
	Email     []string `json:"email"`
	Institute string   `json:"institute"`
}

type Group struct {
	Name     string    `json:"name"`
	Students []Student `json:"students"`
}

// Какое-то мероприятие (ЛР)
type Task struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// КМ-ы
type Endpoint struct {
	Id    int       `json:"-"`
	Tasks []Task    `json:"tasks"`
	Date  time.Time `json:"date"`
}
