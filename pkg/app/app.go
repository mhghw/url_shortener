package app

import (
	"urlShortener/pkg/app/command"
	"urlShortener/pkg/app/query"
)

type Application struct {
	Commands Command
	Queries  Query
}

type Command struct {
	CreateShortLink command.CreateShortLinkHandler
	Register        command.RegisterHandler
}

type Query struct {
	GetOriginalLink query.GetOriginalLinkHandler
	GetUser         query.GetUserHandler
}
