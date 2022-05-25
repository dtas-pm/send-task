package send

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type User struct {
	Id          int          `json:"-" db:"id"`
	Name        string       `json:"name" binding:"required"`
	UserName    string       `json:"username" binding:"required"`
	Password    string       `json:"password" binding:"required"`
	Disciplines []Discipline `json:"disciplines"`
}

type Discipline struct {
	Id    int    `json:"id" db:"id"`
	Name  string `json:"name" db:"name" binding:"required"`
	Event Event  `json:"endpoints" db:"endpoints"`
	Group string `json:"group" db:"groups"`
}

type Event struct {
	Endpoints []Endpoint `json:"endpoints"`
}

func (d Event) Value() (driver.Value, error) {
	return json.Marshal(d)
}

func (a *Event) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &a)
}
