package models

import "context"

type DB interface {
	CheckDBConnection(context.Context)
	Close()
}
