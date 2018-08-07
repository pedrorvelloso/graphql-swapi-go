package types

import "github.com/graphql-go/graphql"

var FilmType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Filme",
		Fields: graphql.Fields{
			"title": &graphql.Field{
				Type: graphql.String,
			},
			"episode_id": &graphql.Field{
				Type: graphql.Int,
			},
			"opening_crawl": &graphql.Field{
				Type: graphql.String,
			},
			"director": &graphql.Field{
				Type: graphql.String,
			},
			"producer": &graphql.Field{
				Type: graphql.String,
			},
			"release_date": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)
