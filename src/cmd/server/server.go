package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"gitlab.com/alex-user-go/art/pkg/adding"
	"gitlab.com/alex-user-go/art/pkg/deleting"
	"gitlab.com/alex-user-go/art/pkg/http/graphql/graph"
	"gitlab.com/alex-user-go/art/pkg/http/graphql/graph/generated"
	"gitlab.com/alex-user-go/art/pkg/listing"
	"gitlab.com/alex-user-go/art/pkg/storage/sqlite"
	"gitlab.com/alex-user-go/art/pkg/updating"
)

const defaultPort = "8088"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	storageType := "sqlite" //hardcoded for simplicity

	var adder adding.Service
	var lister listing.Repository
	var deleter deleting.Repository
	var updater updating.Repository

	switch storageType{
	case "sqlite":
		//s := new(sqlite.Storage)
		s, err := sqlite.NewStorage()
		if err != nil{
			log.Fatal(err)
		}
		adder = adding.NewService(s)
		lister = listing.NewService(s)	
		deleter = deleting.NewService(s)
		updater = updating.NewService(s)
	}

	resolver := &graph.Resolver{Adder: adder, Lister: lister, Deleter: deleter, Updater: updater}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
