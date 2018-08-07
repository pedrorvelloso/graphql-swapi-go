package main

import (
	"fmt"
	"graphql-swapi-go/types"
	"graphql-swapi-go/utils/request"
	"log"
	"net/http"
	"strconv"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

func main() {

	var queryType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"characters": &graphql.Field{
					Type: graphql.NewList(types.CharacterType),
					Args: graphql.FieldConfigArgument{
						"page": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.Int),
						},
					},
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						result := request.DoGet("people/?page=" + strconv.Itoa(p.Args["page"].(int)))
						r := result["results"]
						return r, nil
					},
				},
				"character": &graphql.Field{
					Type: types.CharacterType,
					Args: graphql.FieldConfigArgument{
						"id": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.Int),
						},
					},
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						result := request.DoGet("people/" + strconv.Itoa(p.Args["id"].(int)))
						r := result
						return r, nil
					},
				},
				"planet": &graphql.Field{
					Type: types.PlanetType,
					Args: graphql.FieldConfigArgument{
						"id": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.Int),
						},
					},
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						result := request.DoGet("planets/" + strconv.Itoa(p.Args["id"].(int)))
						return result, nil
					},
				},
			},
		})

	schemaConfig := graphql.SchemaConfig{Query: queryType}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})

	fmt.Println("🚀  Server running on port 8080")
	http.Handle("/graphql", h)
	http.ListenAndServe(":8080", nil)

}
