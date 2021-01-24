package graph

import (
	"gitlab.com/alex-user-go/art/pkg/adding"
	"gitlab.com/alex-user-go/art/pkg/deleting"
	"gitlab.com/alex-user-go/art/pkg/listing"
	"gitlab.com/alex-user-go/art/pkg/updating"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

//Resolver - resolve dependencies
type Resolver struct{
	Adder adding.Repository
	Lister listing.Repository
	Deleter deleting.Repository
	Updater updating.Repository
}
