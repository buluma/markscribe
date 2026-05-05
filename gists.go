// Package main provides a template-driven markdown generator using GitHub, RSS, and other data sources.
package main

import (
	"context"

	"github.com/shurcooL/githubv4"
)

var gistsQuery struct {
	User struct {
		Login githubv4.String
		Gists struct {
			TotalCount githubv4.Int
			Edges      []struct {
				Cursor githubv4.String
				Node   qlGist
			}
		} `graphql:"gists(first: $count, orderBy: {field: CREATED_AT, direction: DESC})"`
	} `graphql:"user(login:$username)"`
}

func gists(count int) []Gist {
	// fmt.Printf("Finding gists...\n")

	var gists []Gist
	variables := map[string]interface{}{
		"username": githubv4.String(username),
		"count":    githubv4.Int(int32(count)),
	}
	err := gitHubClient.Query(context.Background(), &gistsQuery, variables)
	if err != nil {
		panic(err)
	}

	// fmt.Printf("%+v\n", query)
	for _, v := range gistsQuery.User.Gists.Edges {
		gists = append(gists, gistFromQL(v.Node))
	}

	// fmt.Printf("Found %d gists!\n", len(gists))
	return gists
}

/*
{
  user(login: "buluma") {
    login
    gists(first: 100) {
      totalCount
      edges {
        cursor
        node {
		  name
		  description
		  url
		  createdAt
        }
      }
    }
  }
}
*/
