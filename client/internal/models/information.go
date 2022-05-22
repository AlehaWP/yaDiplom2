package models

type User struct {
	ID       int `json:"-"`
	UUID     string
	Login    string
	Email    string
	Phone    string
	Password string
}

type Account struct {
	ID       int `json:"-"`
	UUID     string
	Login    string
	Password string
	User     User `json:"-"`
}

type File struct {
	ID   int `json:"-"`
	UUID string
	Name string
	Data []byte
	User User `json:"-"`
}

type Card struct {
	ID     int `json:"-"`
	UUID   string
	Number string
	Owner  string
	Month  int
	Year   int
	User   User `json:"-"`
}
