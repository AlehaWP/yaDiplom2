package models

type Account struct {
	ID       int
	Login    string
	Password string
}

type File struct {
	ID   int
	Name string
	Data []byte
}

type Card struct {
	ID     int
	Number string
	Owner  string
	Month  int
	Year   int
}
