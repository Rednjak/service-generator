package app

import (
	"zlatane/pkg/app/command"
	"zlatane/pkg/app/query"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

// Commands are used to make changes in the system
type Commands struct {
	CancelResource command.CancelResourceHandler
	CreateResource command.CreateResourceHandler
}

// Queries are used to read the data from the system
type Queries struct {
	AllResources query.AllResourcesHandler
	GetResource  query.GetResourceHandler
}
