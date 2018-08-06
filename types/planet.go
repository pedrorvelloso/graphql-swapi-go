package types

import (
	"graphql-swapi-go/utils/request"

	"github.com/graphql-go/graphql"
)

var PlanetType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Planet",
		Fields: graphql.Fields{
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"rotation_period": &graphql.Field{
				Type: graphql.Int,
			},
			"orbital_period": &graphql.Field{
				Type: graphql.Int,
			},
			"diameter": &graphql.Field{
				Type: graphql.Int,
			},
			"climate": &graphql.Field{
				Type: graphql.String,
			},
			"gravity": &graphql.Field{
				Type: graphql.String,
			},
			"terrain": &graphql.Field{
				Type: graphql.String,
			},
			"surface_water": &graphql.Field{
				Type: graphql.String,
			},
			"population": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

func init() {
	PlanetType.AddFieldConfig("residents", &graphql.Field{
		Type: graphql.NewList(CharacterType),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			source := p.Source.(map[string]interface{})
			residentsLink := source["residents"].([]interface{})
			var results []interface{}
			for _, v := range residentsLink {
				result := request.DoGetFullLink(v.(string))
				results = append(results, result)
			}
			return results, nil
		},
	})
}
