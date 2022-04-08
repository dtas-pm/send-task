package send

type User struct {
	Id          int          `json:"-"`
	Name        string       `json:"name"`
	UserName    string       `json:"username"`
	Password    string       `json:"password"`
	Disciplines []Discipline `json:"disciplines"`
}

type Discipline struct {
	Id        int        `json:"-"`
	Name      string     `json:"name"`
	Endpoints []Endpoint `json:"endpoints"`
	Groups    []Group    `json:"groups"`
}
