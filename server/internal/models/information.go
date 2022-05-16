package models

type User struct {
	ID       int
	UUID     string
	Login    string
	Email    string
	Password string
}

type Account struct {
	ID       int
	UUID     string
	Login    string
	Password string
	User     User
}

type File struct {
	ID   int
	UUID string
	Name string
	Data []byte
	User User
}

type Card struct {
	ID     int
	UUID   string
	Number string
	Owner  string
	Month  int
	Year   int
	User   User
}
