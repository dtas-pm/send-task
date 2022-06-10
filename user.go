package send

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"strings"
	"time"
)

type User struct {
	Id          int          `json:"-" db:"id"`
	Name        string       `json:"name" binding:"required"`
	UserName    string       `json:"username" binding:"required"`
	Password    string       `json:"password" binding:"required"`
	Email       string       `json:"email" db:"email" binding:"required"`
	Role        string       `json:"role" db:"role" binding:"required" `
	Disciplines []Discipline `json:"disciplines"`
}

type Discipline struct {
	Id    int    `json:"id" db:"id"`
	Name  string `json:"name" db:"name" binding:"required"`
	Event Event  `json:"event" db:"endpoints"`
	Group string `json:"group" db:"groups"`
}

type PlanDiscipline struct {
	Id        int       `json:"id" db:"id"`
	Name      string    `json:"name" db:"name" binding:"required"`
	DateStart time.Time `json:"date_start" db:"date_start"`
	Event     EventPD   `json:"event" db:"endpoints"`
	Group     string    `json:"group" db:"groups"`
}

type EventPD struct {
	EndpointsPD []EndpointPD `json:"endpoints_pd"`
}

type EndpointPD struct {
	Name  string `json:"name"`
	Tasks string `json:"tasks"`
	Date  JSTime `json:"date"`
}

type Event struct {
	Endpoints []Endpoint `json:"endpoints"`
}

func (a Event) Value() (driver.Value, error) {
	return json.Marshal(a)
}

func (a *Event) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &a)
}

func (b EventPD) Value() (driver.Value, error) {
	return json.Marshal(b)
}

func (b *EventPD) Scan(value interface{}) error {
	a, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(a, &b)
}

type JSTime time.Time

func (c *JSTime) UnmarshalJSON(b []byte) error {
	value := strings.Trim(string(b), `"`) //get rid of "
	if value == "" || value == "null" {
		return nil
	}

	t, err := time.Parse("2006-01-02", value) //parse time
	if err != nil {
		return err
	}
	*c = JSTime(t) //set result using the pointer
	return nil
}

func (c JSTime) MarshalJSON() ([]byte, error) {
	return []byte(`"` + time.Time(c).Format("2006-01-02") + `"`), nil
}
