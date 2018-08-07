package types

import (
	"graphql-swapi-go/utils/request"

	"github.com/graphql-go/graphql"
)

var SpecieType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Specie",
		Fields: graphql.Fields{
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"classification": &graphql.Field{
				Type: graphql.String,
			},
			"designation": &graphql.Field{
				Type: graphql.String,
			},
			"average_height": &graphql.Field{
				Type: graphql.Int,
			},
			"skin_colors": &graphql.Field{
				Type: graphql.String,
			},
			"hair_colors": &graphql.Field{
				Type: graphql.String,
			},
			"eye_colors": &graphql.Field{
				Type: graphql.String,
			},
			"average_lifespan": &graphql.Field{
				Type: graphql.Int,
			},
			"language": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

func init() {
	SpecieType.AddFieldConfig("homeworld", &graphql.Field{
		Type: PlanetType,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			source := p.Source.(map[string]interface{})
			result := request.DoGetFullLink(source["homeworld"].(string))
			return result, nil
		},
	})
}
