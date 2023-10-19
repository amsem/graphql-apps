package main

import (
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

type Player struct {
    ID int `json:"id"`
    HighScore int `json:"highscore"`
    Name string `json:"name"`
    IsOnline bool `json:"isonline"`
    Location string `json:"location"`
    LevelsUnlocked []string `json:"levelsunlocked"`
}

var players = []Player{
        Player{ID: 123, Name: "amine", HighScore: 80, IsOnline: true, Location: "NY"},
        Player{ID: 432, Name: "said", HighScore: 69, IsOnline: false, Location: "IL"},
}

var playerObject = graphql.NewObject(
    graphql.ObjectConfig{
        Name: "Player",
        Fields: graphql.Fields{
            "id": &graphql.Field{
                Type: graphql.Int,
            },
            "name": &graphql.Field{
                Type: graphql.String,
            },
            "highscore": &graphql.Field{
                Type: graphql.Int,
            },
            "isonline": &graphql.Field{
                Type: graphql.Boolean,
            },
            "location": &graphql.Field{
                Type: graphql.String,
            },
            "levelsunlocked": &graphql.Field{
                Type: graphql.NewList(graphql.String),
            },
        },
    },
)

func main()  {
    fields := graphql.Fields{
        "players": &graphql.Field{
            Type: graphql.NewList(playerObject),
            Description: "All players",
            Resolve: func(p graphql.ResolveParams) (interface{}, error) {
                return players, nil
            },
        },
    }    
    rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
    schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
    schema, _ := graphql.NewSchema(schemaConfig)
    h := handler.New(&handler.Config{
        Schema: &schema,
        Pretty: true,
        GraphiQL: true,
    })
    http.Handle("/graphql", h)
    log.Fatal(http.ListenAndServe(":8000", nil))
}
