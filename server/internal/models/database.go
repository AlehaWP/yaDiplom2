package models

import "context"

type URLs struct {
	ShortURL    string `json:"short_url"`
	OriginalURL string `json:"original_url"`
}

type Statistics struct {
	Urls  string `json:"urls"`
	Users string `json:"users"`
}

//Repository interface repo urls.
type Database interface {
	GetURL(context.Context, string) (string, error)
	SaveURL(context.Context, string, string, string) (string, error)
	SaveURLs(context.Context, map[string]string, string, string) (map[string]string, error)
	FindUser(context.Context, string) bool
	CreateUser(context.Context) (string, error)
	GetUserURLs(context.Context, string) ([]URLs, error)
	CheckDBConnection(context.Context) error
	SetURLsToDel(context.Context, []string, string) error
	GetStatistics(context.Context) (Statistics, error)
}