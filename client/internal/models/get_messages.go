package models

type GetUser struct {
	UUID     string
	Login    string
	Email    string
	Password string
}

type GetAccount struct {
	UUID     string
	Login    string
	Password string
}

type GetFile struct {
	UUID string
	Name string
	Data []byte
}

type GetCard struct {
	UUID   string
	Number string
	Owner  string
	Month  int
	Year   int
}
