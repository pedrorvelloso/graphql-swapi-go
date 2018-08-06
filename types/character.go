package types

import (
	"graphql-swapi-go/utils/request"

	"github.com/graphql-go/graphql"
)

var CharacterType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Character",
		Fields: graphql.Fields{
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"eye_color": &graphql.Field{
				Type: graphql.String,
			},
			"height": &graphql.Field{
				Type: graphql.Int,
			},
			"mass": &graphql.Field{
				Type: graphql.Int,
			},
			"hair_color": &graphql.Field{
				Type: graphql.String,
			},
			"skin_color": &graphql.Field{
				Type: graphql.String,
			},
			"birth_year": &graphql.Field{
				Type: graphql.String,
			},
			"gender": &graphql.Field{
				Type: graphql.String,
			},
			"homeworld": &graphql.Field{
				Type: PlanetType,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					source := p.Source.(map[string]interface{})
					result := request.DoGetFullLink(source["homeworld"].(string))
					return result, nil
				},
			},
		},
	},
)
