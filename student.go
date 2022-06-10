package send

import "time"

type Student struct {
	Id        int      `json:"id" db:"id"`
	FullName  string   `json:"full_name" db:"fullname"`
	Login     string   `json:"login" db:"login"`
	Email     []string `json:"email" db:"email"`
	Institute string   `json:"institute" db:"institute"`
	Group     string   `json:"group" db:"student_group"`
}

type Group struct {
	Name string `json:"name" db:"name"`
}

// Какое-то мероприятие (ЛР)
type Task struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// КМ-ы
type Endpoint struct {
	Id    int       `json:"-"`
	Name  string    `json:"name"`
	Tasks []Task    `json:"tasks"`
	Date  time.Time `json:"date"`
}
